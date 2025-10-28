package core

// ChangeType은 변경 유형
type ChangeType int

const (
	ChangeTypeInsert ChangeType = iota
	ChangeTypeUpdate
	ChangeTypeDelete
)

func (ct ChangeType) String() string {
	return [...]string{"INSERT", "UPDATE", "DELETE"}[ct]
}

// FieldChange는 필드의 변경 정보
type FieldChange struct {
	Field    string // 필드명
	OldValue any    // 이전 값
	NewValue any    // 새 값
}

// Change는 단일 변경사항
type Change struct {
	Type      ChangeType              // 변경 타입
	TableName string                  // 테이블명
	Original  map[string]any          // 원본 값
	Current   map[string]any          // 현재 값
	Delta     map[string]*FieldChange // 변경된 필드만 (UPDATE용)
}

// GetChangedFields는 변경된 필드 목록 반환
func (c *Change) GetChangedFields() []string {
	if c.Delta == nil {
		return []string{}
	}

	fields := make([]string, 0, len(c.Delta))
	for field := range c.Delta {
		fields = append(fields, field)
	}
	return fields
}

// IsEmpty는 변경사항이 비어있는지 확인
func (c *Change) IsEmpty() bool {
	if c.Type == ChangeTypeUpdate {
		return len(c.Delta) == 0
	}
	return false
}
