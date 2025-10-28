package tracking

import (
	"reflect"
	"sync"

	"github.com/homveloper/doodle/delta-outbox-go/core"
)

// ChangeTracker는 엔티티 변경 추적기
type ChangeTracker struct {
	entities map[uintptr]*core.TrackedEntity // 포인터 주소로 추적
	mu       sync.RWMutex
}

// NewChangeTracker는 새로운 변경 추적기 생성
func NewChangeTracker() *ChangeTracker {
	return &ChangeTracker{
		entities: make(map[uintptr]*core.TrackedEntity),
	}
}

// Track은 엔티티 추적 시작
func (ct *ChangeTracker) Track(entity core.Entity, tableName string, state core.EntityState) {
	ct.mu.Lock()
	defer ct.mu.Unlock()

	key := getEntityKey(entity)

	tracked := &core.TrackedEntity{
		Entity:    entity,
		State:     state,
		TableName: tableName,
		Original:  make(map[string]any),
	}

	// 원본 값 스냅샷 저장 (Added가 아닐 때만)
	if state != core.Added {
		tracked.Original = takeSnapshot(entity)
	}

	ct.entities[key] = tracked
}

// GetChanges는 모든 변경사항 반환
func (ct *ChangeTracker) GetChanges() []*core.Change {
	ct.mu.RLock()
	defer ct.mu.RUnlock()

	changes := make([]*core.Change, 0)

	for _, tracked := range ct.entities {
		switch tracked.State {
		case core.Added:
			change := ct.buildInsertChange(tracked)
			changes = append(changes, change)

		case core.Modified:
			change := ct.buildUpdateChange(tracked)
			if change != nil {
				changes = append(changes, change)
			}

		case core.Deleted:
			change := ct.buildDeleteChange(tracked)
			changes = append(changes, change)
		}
	}

	return changes
}

// Clear는 추적 상태 초기화
func (ct *ChangeTracker) Clear() {
	ct.mu.Lock()
	defer ct.mu.Unlock()
	ct.entities = make(map[uintptr]*core.TrackedEntity)
}

// buildInsertChange는 INSERT 변경 생성
func (ct *ChangeTracker) buildInsertChange(tracked *core.TrackedEntity) *core.Change {
	current := getCurrentValues(tracked.Entity)

	return &core.Change{
		Type:      core.ChangeTypeInsert,
		TableName: tracked.TableName,
		Current:   current,
	}
}

// buildUpdateChange는 UPDATE 변경 생성 (델타만)
func (ct *ChangeTracker) buildUpdateChange(tracked *core.TrackedEntity) *core.Change {
	current := getCurrentValues(tracked.Entity)
	delta := make(map[string]*core.FieldChange)

	// 변경된 필드만 추출 (델타 계산)
	for field, currentVal := range current {
		originalVal, exists := tracked.Original[field]

		if !exists || !reflect.DeepEqual(originalVal, currentVal) {
			delta[field] = &core.FieldChange{
				Field:    field,
				OldValue: originalVal,
				NewValue: currentVal,
			}
		}
	}

	// 변경사항이 없으면 nil 반환
	if len(delta) == 0 {
		return nil
	}

	return &core.Change{
		Type:      core.ChangeTypeUpdate,
		TableName: tracked.TableName,
		Original:  tracked.Original,
		Current:   current,
		Delta:     delta,
	}
}

// buildDeleteChange는 DELETE 변경 생성
func (ct *ChangeTracker) buildDeleteChange(tracked *core.TrackedEntity) *core.Change {
	return &core.Change{
		Type:      core.ChangeTypeDelete,
		TableName: tracked.TableName,
		Original:  tracked.Original,
	}
}

// getEntityKey는 엔티티의 고유 키 반환 (포인터 주소)
func getEntityKey(entity core.Entity) uintptr {
	return reflect.ValueOf(entity).Pointer()
}

// takeSnapshot은 엔티티의 현재 값을 스냅샷으로 저장
func takeSnapshot(entity core.Entity) map[string]any {
	snapshot := make(map[string]any)
	val := reflect.ValueOf(entity)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		value := val.Field(i).Interface()
		snapshot[field.Name] = value
	}

	return snapshot
}

// getCurrentValues는 엔티티의 현재 값 반환
func getCurrentValues(entity core.Entity) map[string]any {
	return takeSnapshot(entity)
}
