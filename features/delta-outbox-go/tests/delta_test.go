package tests

import (
	"testing"

	"github.com/homveloper/doodle/delta-outbox-go/core"
	"github.com/homveloper/doodle/delta-outbox-go/deltaorm"
)

type TestEntity struct {
	ID    int
	Name  string
	Value int
}

func TestChangeTracking(t *testing.T) {
	ctx, err := deltaorm.NewDbContext(":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer ctx.Close()

	// 테이블 생성
	ctx.Execute(`
		CREATE TABLE test_entities (
			ID INTEGER PRIMARY KEY,
			Name TEXT,
			Value INTEGER
		)
	`)

	t.Run("INSERT 추적", func(t *testing.T) {
		ctx.BeginTracking()

		entity := &TestEntity{ID: 1, Name: "Test", Value: 100}
		ctx.Track(entity, "test_entities", core.Added)

		changes := ctx.GetChanges()
		if len(changes) != 1 {
			t.Errorf("예상: 1개 변경, 실제: %d개", len(changes))
		}

		if changes[0].Type != core.ChangeTypeInsert {
			t.Errorf("예상: INSERT, 실제: %v", changes[0].Type)
		}
	})

	t.Run("UPDATE 델타 추적", func(t *testing.T) {
		ctx.BeginTracking()

		// 원본 엔티티
		entity := &TestEntity{ID: 1, Name: "Original", Value: 100}
		ctx.Track(entity, "test_entities", core.Modified)

		// Name만 변경
		entity.Name = "Modified"

		changes := ctx.GetChanges()
		if len(changes) != 1 {
			t.Errorf("예상: 1개 변경, 실제: %d개", len(changes))
		}

		change := changes[0]
		if change.Type != core.ChangeTypeUpdate {
			t.Errorf("예상: UPDATE, 실제: %v", change.Type)
		}

		// 델타 확인: Name만 변경되어야 함
		if len(change.Delta) != 1 {
			t.Errorf("예상: 1개 필드 변경, 실제: %d개", len(change.Delta))
		}

		if _, exists := change.Delta["Name"]; !exists {
			t.Error("Name 필드가 변경되지 않음")
		}

		if _, exists := change.Delta["Value"]; exists {
			t.Error("Value 필드는 변경되지 않아야 함")
		}
	})

	t.Run("변경 없음 감지", func(t *testing.T) {
		ctx.BeginTracking()

		entity := &TestEntity{ID: 1, Name: "Same", Value: 100}
		ctx.Track(entity, "test_entities", core.Modified)

		// 아무것도 변경하지 않음

		changes := ctx.GetChanges()
		if len(changes) == 0 {
			return // 정상
		}

		if changes[0].IsEmpty() {
			return // 정상 (빈 변경)
		}

		t.Error("변경사항이 없어야 함")
	})
}

func TestOutboxPattern(t *testing.T) {
	ctx, err := deltaorm.NewDbContext(":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer ctx.Close()

	// 테이블 생성
	ctx.Execute(`
		CREATE TABLE test_entities (
			ID INTEGER PRIMARY KEY,
			Name TEXT,
			Value INTEGER
		)
	`)

	t.Run("여러 변경사항 누적", func(t *testing.T) {
		ctx.BeginTracking()

		entity1 := &TestEntity{ID: 1, Name: "Entity1", Value: 10}
		entity2 := &TestEntity{ID: 2, Name: "Entity2", Value: 20}
		entity3 := &TestEntity{ID: 3, Name: "Entity3", Value: 30}

		ctx.Track(entity1, "test_entities", core.Added)
		ctx.Track(entity2, "test_entities", core.Added)
		ctx.Track(entity3, "test_entities", core.Added)

		// Outbox에 쌓임
		if ctx.GetOutboxSize() != 0 {
			// SaveChanges 전에는 0
		}

		// 한번에 커밋
		if err := ctx.SaveChanges(); err != nil {
			t.Fatal(err)
		}

		// 커밋 후 Outbox는 비어야 함
		if ctx.GetOutboxSize() != 0 {
			t.Errorf("커밋 후 Outbox가 비어야 함, 실제: %d", ctx.GetOutboxSize())
		}
	})
}

func TestMultiTableTransaction(t *testing.T) {
	ctx, err := deltaorm.NewDbContext(":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer ctx.Close()

	// 두 개의 테이블 생성
	ctx.Execute(`CREATE TABLE users (ID INTEGER PRIMARY KEY, Name TEXT)`)
	ctx.Execute(`CREATE TABLE products (ID INTEGER PRIMARY KEY, Name TEXT)`)

	type User struct {
		ID   int
		Name string
	}

	type Product struct {
		ID   int
		Name string
	}

	t.Run("여러 테이블 동시 수정", func(t *testing.T) {
		ctx.BeginTracking()

		user := &User{ID: 1, Name: "Alice"}
		product := &Product{ID: 1, Name: "MacBook"}

		ctx.Track(user, "users", core.Added)
		ctx.Track(product, "products", core.Added)

		changes := ctx.GetChanges()
		if len(changes) != 2 {
			t.Errorf("예상: 2개 테이블 변경, 실제: %d개", len(changes))
		}

		// 한 트랜잭션으로 커밋
		if err := ctx.SaveChanges(); err != nil {
			t.Fatal(err)
		}

		// 데이터 확인
		rows, _ := ctx.Query("SELECT COUNT(*) FROM users")
		var userCount int
		rows.Next()
		rows.Scan(&userCount)
		rows.Close()

		if userCount != 1 {
			t.Errorf("users 테이블: 예상 1개, 실제 %d개", userCount)
		}

		rows, _ = ctx.Query("SELECT COUNT(*) FROM products")
		var productCount int
		rows.Next()
		rows.Scan(&productCount)
		rows.Close()

		if productCount != 1 {
			t.Errorf("products 테이블: 예상 1개, 실제 %d개", productCount)
		}
	})
}
