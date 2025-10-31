# ASCII Art Generator (Go)

텍스트를 멋진 ASCII 아트로 변환하는 Go 패키지입니다.

```
 #   #  #####  #      #      #####
 #   #  #      #      #      #   #
 #####  #####  #      #      #   #
 #   #  #      #      #      #   #
 #   #  #####  #####  #####  #####
```

## 특징

- **간단한 API**: 단일 함수로 모든 기능 제공
- **유연한 커스터마이징**: 옵션 패턴으로 다양한 스타일 지원
- **제로 의존성**: 표준 라이브러리만 사용
- **다양한 폰트**: 여러 ASCII 폰트 스타일
- **Go답게**: Idiomatic Go 패턴 사용

## 빠른 시작

### 설치

```bash
go get github.com/homveloper/doodle/features/ascii-art-go/asciiart
```

### 기본 사용

```go
package main

import (
    "fmt"
    "github.com/homveloper/doodle/features/ascii-art-go/asciiart"
)

func main() {
    result, _ := asciiart.Generate("HELLO")
    fmt.Println(result)
}
```

### 옵션 사용

```go
// 큰 폰트 + 테두리
result, _ := asciiart.Generate("GO",
    asciiart.WithFont(asciiart.FontBig),
    asciiart.WithBorder(),
)

// 모든 옵션 활용
result, _ := asciiart.Generate("DOODLE",
    asciiart.WithFont(asciiart.FontBlock),
    asciiart.WithStyle(asciiart.StyleShadow),
    asciiart.WithAlignment(asciiart.AlignCenter),
    asciiart.WithWidth(80),
    asciiart.WithPadding(2),
    asciiart.WithBorder(),
)
```

## API 문서

### 메인 함수

```go
func Generate(text string, opts ...Option) (string, error)
```

텍스트를 ASCII 아트로 변환합니다.

### 옵션

#### 폰트 선택

```go
WithFont(font Font) Option
```

사용 가능한 폰트:
- `FontStandard` - 기본 5줄 높이 (기본값)
- `FontBig` - 크고 두꺼운 스타일
- `FontSmall` - 작은 3줄 높이
- `FontBlock` - 블록 스타일
- `FontBanner` - 배너 스타일

#### 스타일

```go
WithStyle(style Style) Option
```

사용 가능한 스타일:
- `StyleNormal` - 기본 (기본값)
- `StyleShadow` - 그림자 효과
- `StyleDouble` - 이중선
- `StyleDotted` - 점선 스타일

#### 정렬

```go
WithAlignment(align Align) Option
```

- `AlignLeft` - 왼쪽 정렬 (기본값)
- `AlignCenter` - 가운데 정렬
- `AlignRight` - 오른쪽 정렬

#### 기타 옵션

```go
WithWidth(width int) Option        // 최대 너비 설정 (줄바꿈)
WithPadding(padding int) Option    // 좌우 여백
WithBorder() Option                // 테두리 추가
```

## 예제

더 많은 예제는 [examples/](examples/) 폴더를 참고하세요.

### 기본 생성

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

### 테두리 추가

```go
result, _ := asciiart.Generate("GO",
    asciiart.WithBorder(),
)
fmt.Println(result)
```

출력:
```
╔══════════════════╗
║  ####    #####   ║
║ #    #  #     #  ║
║ #       #     #  ║
║ #  ###  #     #  ║
║ #    #  #     #  ║
║  ####    #####   ║
╚══════════════════╝
```

### 중앙 정렬

```go
result, _ := asciiart.Generate("Hi",
    asciiart.WithAlignment(asciiart.AlignCenter),
    asciiart.WithWidth(30),
)
```

## 개발

### 테스트 실행

```bash
cd asciiart
go test -v
```

### 벤치마크

```bash
go test -bench=.
```

### 예제 실행

```bash
cd examples
go run main.go
```

## 구현 상태

### Phase 1: MVP ✅ (구현 예정)
- [ ] 기본 Generate 함수
- [ ] Standard 폰트 (A-Z, 0-9)
- [ ] 옵션 패턴 구조
- [ ] 기본 테스트

### Phase 2: 확장 (계획 중)
- [ ] 추가 폰트 (Big, Small, Block)
- [ ] 테두리 기능
- [ ] 정렬 기능
- [ ] 패딩 및 너비 제한

### Phase 3: 고급 (향후)
- [ ] 스타일 시스템
- [ ] 소문자 지원
- [ ] 특수문자 지원
- [ ] 여러 줄 처리

## 설계 문서

자세한 설계 내용은 [DESIGN.md](DESIGN.md)를 참고하세요.

## 라이선스

MIT License - 자유롭게 사용하세요!

## 기여

이슈와 PR 환영합니다!

---

**Language**: Go
**Purpose**: ASCII Art 생성 라이브러리
**Status**: 🚧 구현 중
**Feature**: `ascii-art-go`
