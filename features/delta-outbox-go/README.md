# Delta Outbox ORM (Go)

**언어**: Go 1.21+
**목적**: Outbox Pattern + Delta Update를 결합한 경량 ORM

## 📋 개요

Entity Framework의 Change Tracking과 Outbox Pattern을 결합한 Go ORM 라이브러리입니다.
**작업 생애주기 동안 발생한 모든 엔티티 변경사항을 추적**하고, **최소한의 네트워크 트래픽**으로 데이터베이스에 반영합니다.

### 핵심 특징

🎯 **Unit of Work Pattern**
- 작업 단위로 변경사항 관리
- 여러 테이블에 대한 수정을 한 트랜잭션으로 처리

🔍 **Change Tracking (변경 추적)**
- 엔티티의 원본 스냅샷 저장
- 수정 시 변경된 필드만 감지 (Dirty Checking)
- 불필요한 업데이트 제거

⚡ **Delta Update (델타 업데이트)**
- 변경된 필드만 UPDATE 쿼리에 포함
- 네트워크 트래픽 최소화
- 성능 최적화

📦 **Outbox Pattern**
- 변경사항을 메모리에 누적
- 한 번의 트랜잭션으로 일괄 커밋
- 트랜잭션 일관성 보장

## 🏗️ 아키텍처

```
delta-outbox-go/
├── core/                    # 핵심 타입 정의
│   ├── entity.go           # Entity, EntityState
│   └── change.go           # Change, ChangeType
│
├── tracking/                # 변경 추적
│   └── change_tracker.go   # ChangeTracker (스냅샷, 델타 계산)
│
├── outbox/                  # Outbox 패턴
│   └── outbox.go           # Outbox (변경사항 누적, 플러시)
│
├── deltaorm/                # 통합 레이어
│   └── context.go          # DbContext (IoC Container + UoW)
│
├── examples/                # 사용 예제
│   └── main.go
│
└── tests/                   # 테스트
    └── delta_test.go
```

## 🚀 사용 방법

### 기본 사용법

```go
package main

import (
    "github.com/homveloper/doodle/delta-outbox-go/core"
    "github.com/homveloper/doodle/delta-outbox-go/deltaorm"
)

type User struct {
    ID    int
    Name  string
    Email string
    Age   int
}

func main() {
    // 1. DbContext 생성 (IoC Container)
    ctx, _ := deltaorm.NewDbContext("app.db")
    defer ctx.Close()

    // 2. Unit of Work 시작
    ctx.BeginTracking()

    // 3. 엔티티 추가 (INSERT)
    user := &User{ID: 1, Name: "Alice", Email: "alice@example.com", Age: 25}
    ctx.Track(user, "users", core.Added)

    // 4. 커밋
    ctx.SaveChanges()
    // 실행: INSERT INTO users (ID, Name, Email, Age) VALUES (?, ?, ?, ?)
}
```

### 델타 업데이트 (핵심!)

```go
// 1. Unit of Work 시작
ctx.BeginTracking()

// 2. 기존 엔티티 로드 및 추적
user := &User{ID: 1, Name: "Alice", Email: "alice@example.com", Age: 25}
ctx.Track(user, "users", core.Modified)

// 3. 일부 필드만 수정
user.Age = 26  // Age만 변경

// 4. 변경사항 확인
changes := ctx.GetChanges()
fmt.Println(changes[0].GetChangedFields()) // ["Age"]

// 5. 커밋 - 변경된 필드만 업데이트!
ctx.SaveChanges()
// 실행: UPDATE users SET Age = ? WHERE ID = ?
//      (Name, Email은 포함되지 않음!)
```

### 여러 테이블 동시 수정 (Outbox Pattern)

```go
ctx.BeginTracking()

// 여러 엔티티 수정
user := &User{ID: 1, Name: "Alice", Age: 25}
product := &Product{ID: 1, Name: "MacBook", Stock: 10}

ctx.Track(user, "users", core.Modified)
ctx.Track(product, "products", core.Modified)

user.Age = 26
product.Stock = 9

// 한 트랜잭션으로 모든 변경사항 커밋
ctx.SaveChanges()
// BEGIN TRANSACTION
//   UPDATE users SET Age = 26 WHERE ID = 1
//   UPDATE products SET Stock = 9 WHERE ID = 1
// COMMIT
```

## 🧪 테스트

```bash
cd features/delta-outbox-go
go test ./tests/... -v
```

**테스트 커버리지:**
- ✅ INSERT 추적
- ✅ UPDATE 델타 추적
- ✅ 변경 없음 감지
- ✅ Outbox 패턴 (여러 변경사항 누적)
- ✅ 멀티 테이블 트랜잭션

## 📊 실행 예제

```bash
go run examples/main.go
```

**출력:**
```
=== Delta Outbox Pattern Demo ===

📝 추적 시작: 4개 엔티티 추가
   - User: Alice, Bob
   - Product: MacBook, iPhone

✅ SaveChanges #1: 4개 INSERT 실행

📝 수정 작업:
   - User #1: Age 25 → 26
   - Product #1: Stock 10 → 9

🔍 감지된 변경사항: 2개
   [1] UPDATE users: 변경 필드 = [Age]
   [2] UPDATE products: 변경 필드 = [Stock]

✅ SaveChanges #2: 2개 UPDATE 실행 (변경 필드만!)
```

## 🎯 핵심 개념

### 1. Unit of Work (작업 단위)

작업의 생애주기 동안 모든 변경사항을 추적하고 한 번에 커밋:

```go
ctx.BeginTracking()  // 작업 시작
// ... 여러 엔티티 수정 ...
ctx.SaveChanges()    // 한 번에 커밋
```

### 2. Change Tracking (변경 추적)

엔티티의 원본 스냅샷을 저장하고 변경 감지:

```go
// 원본 스냅샷
original := {Name: "Alice", Age: 25}

// 수정
entity.Age = 26

// 델타 계산
delta := {Age: 25 → 26}  // Name은 변경 안됨
```

### 3. Outbox Pattern

변경사항을 메모리에 누적하고 한 번에 플러시:

```
[ChangeTracker] → [Outbox] → [Database]
   변경 감지      누적        일괄 커밋
```

## 💡 실제 사용 시나리오

### 시나리오 1: 주문 처리

```go
ctx.BeginTracking()

// 주문 생성
order := &Order{ID: 1, Status: "Pending", Total: 1000}
ctx.Track(order, "orders", core.Added)

// 재고 감소
product := &Product{ID: 1, Stock: 100}
ctx.Track(product, "products", core.Modified)
product.Stock = 99

// 사용자 포인트 차감
user := &User{ID: 1, Points: 500}
ctx.Track(user, "users", core.Modified)
user.Points = 400

// 한 트랜잭션으로 모든 변경사항 커밋
ctx.SaveChanges()
// Orders, Products, Users 테이블이 원자적으로 업데이트
```

### 시나리오 2: 대량 데이터 업데이트

```go
ctx.BeginTracking()

for _, user := range users {
    ctx.Track(user, "users", core.Modified)
    user.Status = "Active"  // 한 필드만 변경
}

// 변경된 필드만 업데이트 (네트워크 효율적)
ctx.SaveChanges()
// UPDATE users SET Status = 'Active' WHERE ID = ?
// (다른 필드는 전송하지 않음!)
```

## 🔧 기술 스택

- **Go 1.21+**: 제네릭, 리플렉션
- **SQLite3**: 임베디드 데이터베이스
- **Reflection**: 런타임 엔티티 분석
- **Patterns**: Unit of Work, Repository, IoC, Outbox

## 📈 성능 이점

| 전통적인 ORM | Delta Outbox ORM | 개선 |
|-------------|------------------|------|
| 모든 필드 UPDATE | 변경 필드만 UPDATE | 🚀 네트워크 트래픽 감소 |
| 개별 쿼리 실행 | 일괄 트랜잭션 | 🚀 데이터베이스 왕복 감소 |
| 자동 커밋 | 명시적 Unit of Work | 🎯 트랜잭션 제어 향상 |

## 🚧 향후 계획

- [ ] Struct 태그 지원 (`db:"column_name"`)
- [ ] 자동 마이그레이션
- [ ] 관계 매핑 (1:N, N:M)
- [ ] 쿼리 빌더
- [ ] 낙관적 동시성 제어 (Optimistic Concurrency)
- [ ] 배치 INSERT 최적화

## 🎓 배운 점

- **Unit of Work 패턴**: 트랜잭션 경계 명확화
- **Change Tracking**: Entity Framework 핵심 메커니즘
- **Outbox Pattern**: 이벤트 소싱과 마이크로서비스에 유용
- **Go Reflection**: 런타임 타입 분석
- **델타 계산**: 성능 최적화의 핵심

## 📝 라이선스

MIT License - 자유롭게 학습하고 실험하세요!
