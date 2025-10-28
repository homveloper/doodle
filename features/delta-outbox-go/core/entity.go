package core

// Entity는 추적 가능한 엔티티의 기본 인터페이스
type Entity interface{}

// EntityState는 엔티티의 상태를 나타냄
type EntityState int

const (
	Unchanged EntityState = iota // 변경 없음
	Added                         // 새로 추가됨
	Modified                      // 수정됨
	Deleted                       // 삭제됨
)

func (s EntityState) String() string {
	return [...]string{"Unchanged", "Added", "Modified", "Deleted"}[s]
}

// TrackedEntity는 추적 중인 엔티티 정보
type TrackedEntity struct {
	Entity    Entity         // 실제 엔티티
	State     EntityState    // 현재 상태
	Original  map[string]any // 원본 값 (스냅샷)
	TableName string         // 테이블 명
}
