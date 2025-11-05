# Shop - 모바일 이커머스 앱

**Templ + HTMX + Go로 구현한 모바일 최적화 쇼핑몰**

## 소개

Shop은 모바일 퍼스트 디자인으로 구현된 간단하면서도 완전한 기능의 이커머스 웹 애플리케이션입니다. Templ 템플릿 엔진과 HTMX를 사용하여 JavaScript 없이도 동적인 사용자 경험을 제공합니다.

## 주요 기능

### 제품 관리
- 📦 제품 목록 그리드 뷰 (모바일 최적화)
- 🔍 실시간 제품 검색 (HTMX)
- 🏷️ 카테고리별 필터링
- 💰 가격 및 재고 표시

### 장바구니
- 🛒 슬라이드인 장바구니 드로어
- ➕ 수량 조절 (재고 제한 포함)
- 🗑️ 제품 삭제
- 💵 실시간 총액 계산
- 🔔 OOB (Out-of-Band) 배지 업데이트

### 모바일 UX
- 📱 430px 최대 너비 (모바일 중심)
- 👆 터치 친화적 버튼 (최소 44x44px)
- 🏠 하단 네비게이션 바
- 🎨 iOS 스타일 디자인
- ⚡ 빠른 인터랙션

## 기술 스택

```
Backend:     Go 1.21+
Templates:   Templ (Type-safe HTML)
Frontend:    HTMX 1.9+ (Zero JavaScript)
Styling:     Embedded CSS (Mobile-first)
Data:        In-memory store
Testing:     Go standard testing
Dev Tools:   Air (Hot Reload), Makefile
```

## 프로젝트 구조

```
shop-templ/
├── models/              # 데이터 모델 & 비즈니스 로직
│   ├── product.go       # Product 구조체 & 스토어
│   ├── product_test.go  # Product 테스트
│   ├── cart.go          # Cart 로직
│   └── cart_test.go     # Cart 테스트
├── handlers/            # HTTP 핸들러
│   ├── products.go      # 제품 라우트
│   └── cart.go          # 장바구니 라우트
├── templates/           # Templ 컴포넌트
│   ├── layout.templ     # 기본 레이아웃
│   ├── products.templ   # 제품 컴포넌트
│   ├── cart.templ       # 장바구니 컴포넌트
│   └── shared.templ     # 공통 컴포넌트
├── main.go              # 애플리케이션 진입점
└── README.md
```

## 빠른 시작 (Quick Start)

### Makefile 사용 (권장)

```bash
# 1. 프로젝트 초기화 (최초 1회)
make init

# 2. 개발 서버 시작 (Hot Reload)
make dev

# 3. 브라우저에서 열기
# http://localhost:8080
```

### 사용 가능한 명령어

#### 개발
```bash
make dev          # 개발 서버 시작 (Hot Reload)
make dev-clean    # 정리 후 개발 서버 시작
make dev-info     # 개발 서버 정보 표시
make watch        # Watch 모드 정보
```

#### 빌드 & 실행
```bash
make build        # 프로덕션 빌드
make run          # 빌드 + 실행
make clean        # 빌드 아티팩트 정리
```

#### 테스트
```bash
make test         # 테스트 실행
make test-cover   # 테스트 + 커버리지 HTML
```

#### 코드 품질
```bash
make fmt          # 코드 포맷팅 (Go + Templ)
make lint         # 린터 실행
```

#### 유틸리티
```bash
make help         # 모든 명령어 보기
make kill-port    # 포트 8080 정리
make logs         # Air 빌드 로그 확인
make check-deps   # 의존성 확인
```

## 설치 및 실행 (수동)

### 1. 의존성 설치

```bash
# Templ CLI 설치
go install github.com/a-h/templ/cmd/templ@latest

# Air (Hot Reload) 설치
go install github.com/air-verse/air@latest

# Go 모듈 의존성 설치
go mod download
```

### 2. 템플릿 생성 및 빌드

```bash
# Templ 템플릿 생성
templ generate

# 빌드
go build -o shop-server

# 실행
./shop-server
```

### 3. 브라우저에서 열기

```
http://localhost:8080
```

**최적의 경험을 위해 브라우저 개발자 도구에서 모바일 뷰포트(430px)로 설정하세요.**

## 개발 워크플로우

### 방법 1: Air를 이용한 Hot Reload (권장)

```bash
# Air 실행 - .go와 .templ 파일 변경 시 자동 리로드
make dev

# 또는
air
```

**장점:**
- ✅ `.go` 및 `.templ` 파일 변경 시 자동 재빌드
- ✅ 자동으로 `templ generate` 실행
- ✅ 단일 명령어로 개발 환경 실행
- ✅ 빠른 피드백 루프

### 방법 2: 수동 워크플로우

```bash
# 터미널 1: Templ 자동 생성
templ generate --watch

# 터미널 2: 서버 실행
go run .
```

### 테스트 실행

```bash
# 모든 테스트 실행
go test ./...

# 커버리지 포함
go test -cover ./...

# 특정 패키지만 테스트
go test ./models -v
```

## API 엔드포인트

### 제품

| 메서드 | 경로 | 설명 |
|--------|------|------|
| GET | `/` | 홈 (전체 제품 목록) |
| GET | `/products?category=전자제품` | 카테고리별 필터링 |
| GET | `/search?q=검색어` | 제품 검색 |
| GET | `/categories` | 카테고리 목록 |

### 장바구니

| 메서드 | 경로 | 설명 |
|--------|------|------|
| GET | `/cart` | 장바구니 드로어 |
| POST | `/cart/add?product_id=1&quantity=2` | 제품 추가 |
| POST | `/cart/update?product_id=1&quantity=3` | 수량 변경 |
| POST | `/cart/remove?product_id=1` | 제품 제거 |
| POST | `/cart/clear` | 장바구니 비우기 |

## HTMX 패턴

### 실시간 검색
```html
<input
    hx-get="/search"
    hx-trigger="keyup changed delay:300ms"
    hx-target="#product-list"
/>
```

### 장바구니 추가 (OOB 업데이트)
```html
<button
    hx-post="/cart/add?product_id=1&quantity=1"
    hx-target="#cart-badge"
    hx-swap="outerHTML"
>
```

### 카테고리 필터
```html
<button
    hx-get="/products?category=전자제품"
    hx-target="#product-list"
>
```

## 테스트 현황

```
✅ Product 모델: 7개 테스트 (100% 커버리지)
✅ Cart 모델: 10개 테스트 (100% 커버리지)
```

### 주요 테스트 케이스

**Product Tests:**
- 스토어 생성 및 초기화
- 제품 추가 및 ID 할당
- ID로 제품 조회
- 전체 제품 목록
- 검색 (이름/설명)
- 카테고리 필터링
- 고유 카테고리 목록

**Cart Tests:**
- 장바구니 생성
- 제품 추가
- 기존 제품 수량 증가
- 수량 업데이트
- 제품 제거
- 장바구니 비우기
- 전체 개수 및 금액 계산

## 샘플 데이터

애플리케이션은 12개의 샘플 제품으로 시작합니다:

- **전자제품** (7개): 무선 이어폰, 스마트워치, USB-C 케이블, 무선 마우스, 블루투스 스피커, 스마트폰 거치대
- **패션** (3개): 백팩, 노트북 파우치, 캔버스 토트백
- **생활용품** (3개): 텀블러, 손목 보호대, LED 데스크 램프

가격대: ₩15,000 ~ ₩299,000

## 아키텍처 특징

### Type-Safe Templates
Templ은 컴파일 타임에 타입 체크를 제공하여 런타임 오류를 방지합니다.

### Zero JavaScript
HTMX를 사용하여 JavaScript 코드 없이도 SPA와 같은 경험을 제공합니다.

### Thread-Safe Operations
모든 스토어 및 장바구니 작업은 `sync.RWMutex`를 사용하여 동시성 안전을 보장합니다.

### Mobile-First Design
- 최대 너비 430px
- 터치 타겟 최소 44x44px
- 하단 네비게이션
- 슬라이드 인터랙션

## 향후 개선 사항

- [ ] 체크아웃 플로우 구현
- [ ] SQLite 영구 저장소
- [ ] 사용자 인증
- [ ] 제품 이미지 업로드
- [ ] 주문 내역
- [ ] 가격 범위 필터
- [ ] 정렬 기능 (가격, 이름, 최신순)
- [ ] 위시리스트
- [ ] 제품 상세 페이지 모달

## 라이선스

Doodle 프로젝트의 일부로, 실험 및 학습 목적으로 자유롭게 사용 가능합니다.

## 기여

이 프로젝트는 Doodle 실험 저장소의 일부입니다. 개선 사항이나 버그 수정은 풀 리퀘스트를 통해 기여해주세요.

---

**Built with ❤️ using Templ, HTMX, and Go**
