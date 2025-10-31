# ASCII Art Generator - 설계 문서

## 📐 프로젝트 개요

Go 언어로 구현하는 ASCII Art Generator 패키지입니다. 텍스트를 입력받아 다양한 스타일의 ASCII 아트로 변환합니다.

**핵심 철학**:
- 단일 함수 진입점
- 옵션 패턴으로 유연한 커스터마이징
- 순수한 라이브러리 (의존성 최소화)

---

## 🏗️ 프로젝트 구조

```
features/ascii-art-go/
├── asciiart/              # 메인 패키지
│   ├── generator.go       # 핵심 생성 로직
│   ├── options.go         # 옵션 패턴 정의
│   ├── fonts.go           # 폰트 데이터
│   └── generator_test.go  # 테스트
├── examples/
│   └── main.go           # 사용 예제
├── DESIGN.md             # 이 문서
├── README.md
└── go.mod
```

---

## 🎯 API 설계

### 메인 함수

```go
// Generate는 주어진 텍스트를 ASCII 아트로 변환합니다.
func Generate(text string, opts ...Option) (string, error)
```

**특징**:
- 단일 진입점 (Single Entry Point)
- 가변 인자로 옵션 전달
- 에러 처리 포함

---

## 🔧 타입 시스템

### 1. Config 구조체

```go
type Config struct {
    Font      Font      // 사용할 폰트
    Width     int       // 최대 너비 (0 = 무제한)
    Padding   int       // 좌우 여백
    Alignment Align     // 정렬 방식
    Border    bool      // 테두리 추가 여부
    Style     Style     // 스타일
}
```

### 2. Option 패턴

```go
type Option func(*Config)

// 옵션 생성자들
func WithFont(font Font) Option
func WithWidth(width int) Option
func WithPadding(padding int) Option
func WithAlignment(align Align) Option
func WithBorder() Option
func WithStyle(style Style) Option
```

**장점**:
- 선택적 파라미터
- 기본값 설정 용이
- 확장성 우수
- 타입 안정성

### 3. 상수 정의

```go
// Font 타입
type Font string

const (
    FontStandard Font = "standard"  // 기본 5줄 높이
    FontBig      Font = "big"       // 크고 두꺼운
    FontSmall    Font = "small"     // 3줄 높이
    FontBlock    Font = "block"     // 블록 스타일
    FontBanner   Font = "banner"    // 배너 스타일
)

// Style 타입
type Style string

const (
    StyleNormal Style = "normal"
    StyleShadow Style = "shadow"   // 그림자 효과
    StyleDouble Style = "double"   // 이중선
    StyleDotted Style = "dotted"   // 점선 스타일
)

// Align 타입
type Align string

const (
    AlignLeft   Align = "left"
    AlignCenter Align = "center"
    AlignRight  Align = "right"
)
```

---

## 🔄 데이터 흐름

```
입력: "HELLO"
    ↓
┌─────────────────────────────┐
│ 1. 옵션 파싱 & Config 생성  │
└─────────────────────────────┘
    ↓
┌─────────────────────────────┐
│ 2. 문자별 폰트 데이터 조회  │
│   'H' → [][]string          │
│   'E' → [][]string          │
│   ...                       │
└─────────────────────────────┘
    ↓
┌─────────────────────────────┐
│ 3. 가로 결합                │
│   (Horizontal Concatenation)│
└─────────────────────────────┘
    ↓
┌─────────────────────────────┐
│ 4. 스타일 적용              │
│   (그림자, 이중선 등)       │
└─────────────────────────────┘
    ↓
┌─────────────────────────────┐
│ 5. 정렬 & 패딩              │
└─────────────────────────────┘
    ↓
┌─────────────────────────────┐
│ 6. 테두리 (옵션)            │
└─────────────────────────────┘
    ↓
출력: ASCII Art String
```

---

## 💾 폰트 데이터 구조

### 내부 표현

```go
type fontData struct {
    height int                      // 문자 높이 (줄 수)
    chars  map[rune][]string        // 문자 → ASCII 패턴
}
```

### 예시: 'A' 문자 (Standard Font)

```go
'A': []string{
    "   ##   ",
    "  #  #  ",
    " #    # ",
    " ###### ",
    " #    # ",
}
```

### 문자 지원 범위

**Phase 1 (MVP)**:
- 대문자: A-Z
- 숫자: 0-9
- 공백

**Phase 2**:
- 소문자: a-z
- 기본 특수문자: ! @ # $ % ^ & * ( ) - _ = +

**Phase 3**:
- 확장 특수문자
- 유니코드 일부 지원 검토

---

## 🎨 핵심 알고리즘

### 1. 문자 가로 결합

```go
func combineHorizontally(chars [][]string) []string {
    if len(chars) == 0 {
        return []string{}
    }

    height := len(chars[0])
    result := make([]string, height)

    for row := 0; row < height; row++ {
        for _, char := range chars {
            result[row] += char[row]
        }
    }

    return result
}
```

### 2. 테두리 추가

```go
func addBorder(lines []string) []string {
    if len(lines) == 0 {
        return lines
    }

    width := len(lines[0])
    result := make([]string, 0, len(lines)+2)

    // 상단
    result = append(result, "╔"+strings.Repeat("═", width)+"╗")

    // 본문
    for _, line := range lines {
        result = append(result, "║"+line+"║")
    }

    // 하단
    result = append(result, "╚"+strings.Repeat("═", width)+"╝")

    return result
}
```

### 3. 정렬

```go
func alignLine(line string, width int, align Align) string {
    lineLen := len(line)
    if lineLen >= width {
        return line
    }

    padding := width - lineLen

    switch align {
    case AlignCenter:
        left := padding / 2
        right := padding - left
        return strings.Repeat(" ", left) + line + strings.Repeat(" ", right)
    case AlignRight:
        return strings.Repeat(" ", padding) + line
    default: // AlignLeft
        return line + strings.Repeat(" ", padding)
    }
}
```

---

## 📊 사용 예시

### 기본 사용

```go
result, _ := asciiart.Generate("HELLO")
fmt.Println(result)
```

출력:
```
 #   #  #####  #      #      #####
 #   #  #      #      #      #   #
 #####  #####  #      #      #   #
 #   #  #      #      #      #   #
 #   #  #####  #####  #####  #####
```

### 옵션 활용

```go
result, _ := asciiart.Generate("GO",
    asciiart.WithFont(asciiart.FontBig),
    asciiart.WithBorder(),
    asciiart.WithAlignment(asciiart.AlignCenter),
)
fmt.Println(result)
```

출력:
```
╔════════════════════════════╗
║    ####    #####           ║
║   #    #  #     #          ║
║   #       #     #          ║
║   #  ###  #     #          ║
║   #    #  #     #          ║
║    ####    #####           ║
╚════════════════════════════╝
```

### 모든 옵션

```go
result, _ := asciiart.Generate("DOODLE",
    asciiart.WithFont(asciiart.FontBlock),
    asciiart.WithStyle(asciiart.StyleShadow),
    asciiart.WithAlignment(asciiart.AlignCenter),
    asciiart.WithWidth(80),
    asciiart.WithPadding(2),
    asciiart.WithBorder(),
)
```

---

## ✅ 구현 단계

### Phase 1: MVP (Minimum Viable Product)

**목표**: 기본 기능 동작

- [ ] 프로젝트 구조 생성
- [ ] `generator.go`: 기본 Generate 함수
- [ ] `options.go`: Option 패턴 구현
- [ ] `fonts.go`: Standard 폰트 1개 (A-Z, 0-9)
- [ ] 기본 테스트
- [ ] 예제 프로그램

**완료 조건**:
```go
asciiart.Generate("HELLO") // 성공적으로 출력
```

### Phase 2: 기능 확장

**목표**: 주요 옵션 구현

- [ ] 추가 폰트 3개 (Big, Small, Block)
- [ ] 테두리 기능
- [ ] 정렬 기능 (Left, Center, Right)
- [ ] 패딩 기능
- [ ] 폭 제한 기능
- [ ] 확장 테스트

**완료 조건**:
```go
asciiart.Generate("GO",
    WithFont(FontBig),
    WithBorder(),
) // 모든 옵션 동작
```

### Phase 3: 고급 기능

**목표**: 스타일과 확장성

- [ ] 스타일 시스템 (Shadow, Double, Dotted)
- [ ] 소문자 지원
- [ ] 특수문자 지원
- [ ] 여러 줄 입력 처리 (`\n`)
- [ ] 성능 최적화
- [ ] 벤치마크 테스트

**완료 조건**:
```go
asciiart.Generate("Hello\nWorld",
    WithStyle(StyleShadow),
) // 고급 기능 모두 동작
```

---

## 🧪 테스트 전략

### 단위 테스트

```go
func TestGenerate(t *testing.T) {
    tests := []struct {
        name    string
        input   string
        opts    []Option
        wantErr bool
        check   func(string) bool
    }{
        {
            name:  "단일 문자",
            input: "A",
            check: func(s string) bool {
                return len(s) > 0 && strings.Contains(s, "#")
            },
        },
        {
            name:  "여러 문자",
            input: "ABC",
            check: func(s string) bool {
                lines := strings.Split(s, "\n")
                return len(lines) >= 3 // 최소 높이
            },
        },
        {
            name:  "빈 문자열",
            input: "",
            check: func(s string) bool {
                return s == ""
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result, err := Generate(tt.input, tt.opts...)
            if (err != nil) != tt.wantErr {
                t.Errorf("unexpected error: %v", err)
            }
            if !tt.check(result) {
                t.Errorf("validation failed for result:\n%s", result)
            }
        })
    }
}
```

### 테스트 항목

**기능 테스트**:
- 단일 문자 생성
- 여러 문자 생성
- 숫자 생성
- 공백 처리
- 빈 문자열 처리
- 미지원 문자 처리

**옵션 테스트**:
- 각 폰트별 테스트
- 테두리 적용
- 정렬 (Left, Center, Right)
- 패딩 적용
- 스타일 적용

**엣지 케이스**:
- 매우 긴 문자열
- 특수문자
- 유니코드
- nil 옵션
- 잘못된 옵션 값

---

## 🎯 성공 지표

### 기능적 목표

- [ ] A-Z, 0-9 완벽 지원
- [ ] 최소 3개 폰트 제공
- [ ] 모든 옵션 정상 동작
- [ ] 테스트 커버리지 80% 이상

### 품질 목표

- [ ] 명확한 API 문서
- [ ] 실행 가능한 예제
- [ ] 직관적인 사용법
- [ ] 제로 의존성 (표준 라이브러리만)

### 성능 목표

- [ ] 100자 변환 < 1ms
- [ ] 메모리 효율적 (불필요한 할당 최소화)
- [ ] Goroutine-safe (동시성 고려)

---

## 🚀 확장 가능성

### 단기 확장

- CLI 도구로 래핑
- 파일 입력/출력 지원
- 컬러 출력 (ANSI 코드)

### 중기 확장

- FIGlet 폰트 파일 로딩
- 커스텀 폰트 제작 도구
- HTTP API 서버

### 장기 확장

- 이미지 → ASCII 변환
- 애니메이션 ASCII
- 웹 프론트엔드

---

## 📚 참고 자료

### Go 옵션 패턴

- [Functional Options Pattern in Go](https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis)
- [Self-referential functions and the design of options](https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html)

### ASCII Art

- [FIGlet](http://www.figlet.org/) - ASCII Art 생성기 원조
- [ASCII Art Archive](https://www.asciiart.eu/)
- [Patorjk's ASCII Generator](http://patorjk.com/software/taag/)

### 테스트

- [Table Driven Tests in Go](https://dave.cheney.net/2019/05/07/prefer-table-driven-tests)
- [Testing Best Practices](https://github.com/golang/go/wiki/TestComments)

---

## 📝 노트

### 설계 결정 사항

1. **왜 옵션 패턴?**
   - Go에는 선택적 파라미터가 없음
   - 빌더 패턴보다 간결
   - 확장성 우수

2. **왜 map[rune][]string?**
   - rune: 유니코드 지원
   - []string: 각 줄을 독립적으로 관리
   - map: O(1) 조회 성능

3. **에러 처리**
   - 미지원 문자: 건너뛰기 or 에러?
   - 초기에는 '?' 같은 플레이스홀더 사용 고려

### 구현 시 주의사항

- UTF-8 인코딩 고려
- 탭/특수 공백 처리
- 줄 길이 일관성 유지
- 메모리 효율 (큰 문자열 처리 시)

---

**작성일**: 2025-10-31
**버전**: 1.0
**상태**: 설계 완료, 구현 대기
