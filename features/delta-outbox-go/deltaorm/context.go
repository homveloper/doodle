package deltaorm

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/homveloper/doodle/delta-outbox-go/core"
	"github.com/homveloper/doodle/delta-outbox-go/outbox"
	"github.com/homveloper/doodle/delta-outbox-go/tracking"
)

// DbContext는 데이터베이스 컨텍스트 (IoC Container + Unit of Work)
type DbContext struct {
	db      *sql.DB
	tracker *tracking.ChangeTracker
	outbox  *outbox.Outbox
}

// NewDbContext는 새로운 DB 컨텍스트 생성
func NewDbContext(dbPath string) (*DbContext, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	ctx := &DbContext{
		db:      db,
		tracker: tracking.NewChangeTracker(),
		outbox:  outbox.NewOutbox(),
	}

	return ctx, nil
}

// Close는 DB 연결 종료
func (ctx *DbContext) Close() error {
	return ctx.db.Close()
}

// BeginTracking은 변경 추적 시작 (Unit of Work 시작)
func (ctx *DbContext) BeginTracking() {
	ctx.tracker.Clear()
	ctx.outbox.Clear()
}

// Track은 엔티티 추적 시작
func (ctx *DbContext) Track(entity core.Entity, tableName string, state core.EntityState) {
	ctx.tracker.Track(entity, tableName, state)
}

// SaveChanges는 모든 변경사항을 커밋 (Unit of Work 커밋)
// 핵심: 여러 테이블의 변경사항을 한 트랜잭션으로!
func (ctx *DbContext) SaveChanges() error {
	// 1. Change Tracker에서 모든 변경사항 수집
	changes := ctx.tracker.GetChanges()

	// 2. Outbox에 추가 (네트워크 최적화)
	for _, change := range changes {
		ctx.outbox.Add(change)
	}

	// 3. Outbox의 변경사항을 한번에 DB에 플러시
	return ctx.outbox.Flush(ctx.db)
}

// GetChanges는 현재 추적 중인 변경사항 조회 (디버깅용)
func (ctx *DbContext) GetChanges() []*core.Change {
	return ctx.tracker.GetChanges()
}

// GetOutboxSize는 Outbox에 쌓인 변경사항 개수
func (ctx *DbContext) GetOutboxSize() int {
	return ctx.outbox.Size()
}

// Execute는 직접 SQL 실행 (헬퍼)
func (ctx *DbContext) Execute(query string, args ...any) error {
	_, err := ctx.db.Exec(query, args...)
	return err
}

// Query는 쿼리 실행 (헬퍼)
func (ctx *DbContext) Query(query string, args ...any) (*sql.Rows, error) {
	return ctx.db.Query(query, args...)
}
