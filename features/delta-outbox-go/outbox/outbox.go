package outbox

import (
	"database/sql"
	"fmt"
	"strings"
	"sync"

	"github.com/homveloper/doodle/delta-outbox-go/core"
)

// Outbox는 변경사항을 모아두는 아웃박스
type Outbox struct {
	changes []*core.Change
	mu      sync.Mutex
}

// NewOutbox는 새로운 Outbox 생성
func NewOutbox() *Outbox {
	return &Outbox{
		changes: make([]*core.Change, 0),
	}
}

// Add는 변경사항 추가
func (o *Outbox) Add(change *core.Change) {
	o.mu.Lock()
	defer o.mu.Unlock()

	// 빈 변경사항은 추가하지 않음
	if change.IsEmpty() {
		return
	}

	o.changes = append(o.changes, change)
}

// Size는 Outbox에 쌓인 변경사항 개수
func (o *Outbox) Size() int {
	o.mu.Lock()
	defer o.mu.Unlock()
	return len(o.changes)
}

// Clear는 Outbox 초기화
func (o *Outbox) Clear() {
	o.mu.Lock()
	defer o.mu.Unlock()
	o.changes = make([]*core.Change, 0)
}

// Flush는 모든 변경사항을 DB에 반영
func (o *Outbox) Flush(db *sql.DB) error {
	o.mu.Lock()
	defer o.mu.Unlock()

	if len(o.changes) == 0 {
		return nil
	}

	// 트랜잭션 시작
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("트랜잭션 시작 실패: %w", err)
	}

	// 모든 변경사항 실행
	for _, change := range o.changes {
		if err := o.executeChange(tx, change); err != nil {
			tx.Rollback()
			return fmt.Errorf("변경사항 실행 실패: %w", err)
		}
	}

	// 커밋
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("커밋 실패: %w", err)
	}

	// 성공 후 초기화
	o.changes = make([]*core.Change, 0)

	return nil
}

// executeChange는 단일 변경사항 실행
func (o *Outbox) executeChange(tx *sql.Tx, change *core.Change) error {
	switch change.Type {
	case core.ChangeTypeInsert:
		return o.executeInsert(tx, change)
	case core.ChangeTypeUpdate:
		return o.executeUpdate(tx, change)
	case core.ChangeTypeDelete:
		return o.executeDelete(tx, change)
	default:
		return fmt.Errorf("알 수 없는 변경 타입: %v", change.Type)
	}
}

// executeInsert는 INSERT 실행
func (o *Outbox) executeInsert(tx *sql.Tx, change *core.Change) error {
	columns := make([]string, 0, len(change.Current))
	placeholders := make([]string, 0, len(change.Current))
	values := make([]any, 0, len(change.Current))

	for col, val := range change.Current {
		columns = append(columns, col)
		placeholders = append(placeholders, "?")
		values = append(values, val)
	}

	query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		change.TableName,
		strings.Join(columns, ", "),
		strings.Join(placeholders, ", "),
	)

	_, err := tx.Exec(query, values...)
	return err
}

// executeUpdate는 UPDATE 실행 (델타만!)
func (o *Outbox) executeUpdate(tx *sql.Tx, change *core.Change) error {
	// 변경된 필드만 업데이트
	setClauses := make([]string, 0, len(change.Delta))
	values := make([]any, 0, len(change.Delta))

	for field, fieldChange := range change.Delta {
		setClauses = append(setClauses, fmt.Sprintf("%s = ?", field))
		values = append(values, fieldChange.NewValue)
	}

	// Primary Key 조건 (ID 필드 사용, 실제로는 메타데이터에서 가져와야 함)
	whereClause := "ID = ?"
	idValue := change.Original["ID"]
	values = append(values, idValue)

	query := fmt.Sprintf(
		"UPDATE %s SET %s WHERE %s",
		change.TableName,
		strings.Join(setClauses, ", "),
		whereClause,
	)

	_, err := tx.Exec(query, values...)
	return err
}

// executeDelete는 DELETE 실행
func (o *Outbox) executeDelete(tx *sql.Tx, change *core.Change) error {
	whereClause := "ID = ?"
	idValue := change.Original["ID"]

	query := fmt.Sprintf(
		"DELETE FROM %s WHERE %s",
		change.TableName,
		whereClause,
	)

	_, err := tx.Exec(query, idValue)
	return err
}

// GetChanges는 현재 Outbox의 변경사항 조회 (디버깅용)
func (o *Outbox) GetChanges() []*core.Change {
	o.mu.Lock()
	defer o.mu.Unlock()
	return append([]*core.Change{}, o.changes...)
}
