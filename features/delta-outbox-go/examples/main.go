package main

import (
	"fmt"
	"log"

	"github.com/homveloper/doodle/delta-outbox-go/core"
	"github.com/homveloper/doodle/delta-outbox-go/deltaorm"
)

// User 엔티티
type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

// Product 엔티티
type Product struct {
	ID    int
	Name  string
	Price int
	Stock int
}

func main() {
	// 1. DbContext 생성 (IoC Container)
	ctx, err := deltaorm.NewDbContext(":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer ctx.Close()

	// 2. 테이블 생성
	setupDatabase(ctx)

	fmt.Println("=== Delta Outbox Pattern Demo ===\n")

	// 3. Unit of Work 시작
	ctx.BeginTracking()

	// 4. 여러 엔티티 생성 및 추적
	user1 := &User{ID: 1, Name: "Alice", Email: "alice@example.com", Age: 25}
	user2 := &User{ID: 2, Name: "Bob", Email: "bob@example.com", Age: 30}
	product1 := &Product{ID: 1, Name: "MacBook", Price: 2000, Stock: 10}
	product2 := &Product{ID: 2, Name: "iPhone", Price: 1000, Stock: 50}

	// Added 상태로 추적
	ctx.Track(user1, "users", core.Added)
	ctx.Track(user2, "users", core.Added)
	ctx.Track(product1, "products", core.Added)
	ctx.Track(product2, "products", core.Added)

	fmt.Println("📝 추적 시작: 4개 엔티티 추가")
	fmt.Printf("   - User: %s, %s\n", user1.Name, user2.Name)
	fmt.Printf("   - Product: %s, %s\n\n", product1.Name, product2.Name)

	// 5. 첫 번째 SaveChanges (4개 INSERT)
	if err := ctx.SaveChanges(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("✅ SaveChanges #1: 4개 INSERT 실행\n")

	// 6. 새로운 Unit of Work 시작
	ctx.BeginTracking()

	// 7. 기존 데이터를 Modified로 추적하고 수정
	user1Loaded := &User{ID: 1, Name: "Alice", Email: "alice@example.com", Age: 25}
	product1Loaded := &Product{ID: 1, Name: "MacBook", Price: 2000, Stock: 10}

	ctx.Track(user1Loaded, "users", core.Modified)
	ctx.Track(product1Loaded, "products", core.Modified)

	// 일부 필드만 수정 (델타 업데이트!)
	user1Loaded.Age = 26 // Age만 변경
	product1Loaded.Stock = 9 // Stock만 변경

	fmt.Println("📝 수정 작업:")
	fmt.Printf("   - User #1: Age 25 → 26\n")
	fmt.Printf("   - Product #1: Stock 10 → 9\n\n")

	// 8. 변경사항 확인
	changes := ctx.GetChanges()
	fmt.Printf("🔍 감지된 변경사항: %d개\n", len(changes))
	for i, change := range changes {
		fmt.Printf("   [%d] %s %s: ", i+1, change.Type, change.TableName)
		if change.Type == core.ChangeTypeUpdate {
			fmt.Printf("변경 필드 = %v\n", change.GetChangedFields())
		} else {
			fmt.Println()
		}
	}
	fmt.Println()

	// 9. 두 번째 SaveChanges (2개 UPDATE, 델타만!)
	if err := ctx.SaveChanges(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("✅ SaveChanges #2: 2개 UPDATE 실행 (변경 필드만!)\n")

	// 10. Outbox 패턴 시연: 여러 테이블 동시 수정
	fmt.Println("=== Outbox Pattern: 여러 테이블 동시 수정 ===\n")

	ctx.BeginTracking()

	// User와 Product를 동시에 수정
	user2Loaded := &User{ID: 2, Name: "Bob", Email: "bob@example.com", Age: 30}
	product2Loaded := &Product{ID: 2, Name: "iPhone", Price: 1000, Stock: 50}

	ctx.Track(user2Loaded, "users", core.Modified)
	ctx.Track(product2Loaded, "products", core.Modified)

	user2Loaded.Name = "Bob Smith"
	user2Loaded.Email = "bob.smith@example.com"
	product2Loaded.Price = 900
	product2Loaded.Stock = 45

	fmt.Printf("📦 Outbox에 쌓인 변경: %d개 (커밋 전)\n", ctx.GetOutboxSize())

	// 11. 한 트랜잭션으로 모든 변경사항 커밋!
	if err := ctx.SaveChanges(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("✅ 한 트랜잭션으로 모든 변경사항 커밋 완료!\n")
	fmt.Println("=== 핵심 이점 ===")
	fmt.Println("✨ 네트워크 트래픽 최소화: 변경 필드만 전송")
	fmt.Println("✨ 트랜잭션 일관성: 여러 테이블 동시 커밋")
	fmt.Println("✨ Unit of Work: 작업 생애주기 동안 변경 추적")
}

func setupDatabase(ctx *deltaorm.DbContext) {
	ctx.Execute(`
		CREATE TABLE IF NOT EXISTS users (
			ID INTEGER PRIMARY KEY,
			Name TEXT,
			Email TEXT,
			Age INTEGER
		)
	`)

	ctx.Execute(`
		CREATE TABLE IF NOT EXISTS products (
			ID INTEGER PRIMARY KEY,
			Name TEXT,
			Price INTEGER,
			Stock INTEGER
		)
	`)
}
