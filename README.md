# 🎨 Doodle

Claude Code Web 실험 프로젝트 - 새로운 기능과 시스템을 TDD로 검증하는 놀이터

## 📋 프로젝트 소개

이 프로젝트는 [Claude Code Web](https://claude.com/claude-code)의 새로운 기능들을 탐색하고 실험하는 공간입니다.

**핵심 철학:**
- 🌐 **언어 독립적**: 각 feature마다 가장 적합한 언어를 선택 (JavaScript, Python, Rust, Go 등)
- 📦 **Feature 중심 구조**: 기능별로 독립적인 폴더에서 관리
- ✅ **TDD 기반 개발**: 모든 코드는 테스트와 함께
- 🚀 **자동화된 검증**: GitHub Actions CI/CD로 다양한 환경에서 테스트

## 🏗️ 프로젝트 구조

Feature 중심의 모듈형 구조로 설계되었습니다:

```
doodle/
├── features/                    # 모든 실험 기능들
│   ├── calculator-js/          # JavaScript 계산기
│   │   ├── calculator.js       # 계산기 로직
│   │   ├── test.js            # 테스트 스위트
│   │   ├── package.json       # Node.js 설정
│   │   └── README.md          # 기능 문서
│   │
│   ├── algorithms-python/      # Python 알고리즘
│   │   ├── test_python.py     # 알고리즘 & 테스트
│   │   └── README.md          # 기능 문서
│   │
│   └── [새로운 features...]   # 앞으로 추가될 실험들
│
├── .github/
│   └── workflows/
│       └── test.yml           # CI/CD 파이프라인
├── CLAUDE.md                  # Claude Code 활용 가이드
├── README.md                  # 이 파일
└── LICENSE
```

### 현재 Features

| Feature | 언어 | 설명 | 상태 |
|---------|------|------|------|
| **calculator-js** | JavaScript | 기본 수학 연산 함수 모음 | ✅ 완료 |
| **algorithms-python** | Python | Fibonacci, 소수 판별 등 | ✅ 완료 |

## 🧪 테스트 주도 개발 (TDD)

각 feature는 독립적인 테스트 스위트를 가지고 있습니다:

### Calculator (JavaScript)
```bash
cd features/calculator-js
npm test
```
**커버리지:** 15개 테스트 - 덧셈, 뺄셈, 곱셈, 나눗셈, 거듭제곱, 모듈로

### Algorithms (Python)
```bash
cd features/algorithms-python
python3 test_python.py
```
**커버리지:** 8개 테스트 - Fibonacci, 소수 판별

### 전체 테스트 실행
모든 features를 한번에 테스트하려면 CI/CD가 자동으로 처리합니다.

## 🚀 CI/CD 파이프라인

GitHub Actions를 통해 모든 푸시와 Pull Request에서 자동으로 테스트를 실행합니다.

### 테스트 환경

**Node.js:**
- Node.js 18.x
- Node.js 20.x
- Node.js 22.x

**Python:**
- Python 3.9
- Python 3.10
- Python 3.11
- Python 3.12

모든 테스트는 병렬로 실행되며, 결과는 GitHub Actions 대시보드에서 확인할 수 있습니다.

## 🎯 사용 방법

### 로컬 개발

1. **저장소 클론**
   ```bash
   git clone https://github.com/homveloper/doodle.git
   cd doodle
   ```

2. **특정 feature 테스트**
   ```bash
   # JavaScript 계산기
   cd features/calculator-js
   npm test

   # Python 알고리즘
   cd features/algorithms-python
   python3 test_python.py
   ```

### 새로운 Feature 추가

새로운 실험을 시작할 때:

1. **언어 선택**: 해당 feature에 가장 적합한 언어 결정
   - 웹 관련? → JavaScript/TypeScript
   - 데이터 처리? → Python
   - 성능 중요? → Rust/Go
   - 학습 목적? → 배우고 싶은 언어

2. **폴더 구조 생성**
   ```bash
   mkdir -p features/my-feature-lang
   cd features/my-feature-lang
   ```

3. **TDD로 개발**
   - 테스트 먼저 작성
   - 구현
   - 테스트 통과 확인

4. **문서화**
   - README.md 작성
   - 언어, 목적, 사용법 명시

5. **CI/CD 자동 실행**
   - Git 커밋 & 푸시
   - GitHub Actions가 자동으로 테스트

## 🤖 Claude Code와 함께

이 프로젝트는 Claude Code Web을 최대한 활용하도록 설계되었습니다. 자세한 내용은 [CLAUDE.md](./CLAUDE.md)를 참고하세요.

## 📊 테스트 상태

![Tests](https://github.com/homveloper/doodle/actions/workflows/test.yml/badge.svg)

## 📝 라이선스

MIT License - 자유롭게 실험하고 배우세요!

## 🔗 참고 자료

- [Claude Code 문서](https://docs.claude.com/claude-code)
- [GitHub Actions 문서](https://docs.github.com/actions)
- [TDD 가이드](https://en.wikipedia.org/wiki/Test-driven_development)
