# 🎨 Doodle

Claude Code Web 실험 프로젝트 - 새로운 기능과 시스템을 TDD로 검증하는 놀이터

## 📋 프로젝트 소개

이 프로젝트는 [Claude Code Web](https://claude.com/claude-code)의 새로운 기능들을 탐색하고 실험하는 공간입니다. 다양한 기능, 시스템, 코드를 폴더별로 구분하여 시도해보고, TDD(Test-Driven Development)와 GitHub Actions CI/CD를 통해 자동으로 테스트합니다.

## 🏗️ 프로젝트 구조

현재 프로젝트는 다음과 같은 실험들을 포함하고 있습니다:

```
doodle/
├── calculator.js          # Node.js 계산기 예제
├── test.js               # Node.js 유닛 테스트
├── test_python.py        # Python 테스트 (Fibonacci, Prime)
├── .github/
│   └── workflows/
│       └── test.yml      # CI/CD 파이프라인
├── package.json          # Node.js 프로젝트 설정
└── README.md            # 이 파일
```

## 🧪 테스트 주도 개발 (TDD)

### Node.js 테스트

프로젝트는 외부 의존성 없이 자체 테스트 러너를 사용합니다:

```bash
npm test
```

**테스트 커버리지:**
- ✓ 덧셈, 뺄셈, 곱셈, 나눗셈
- ✓ 거듭제곱, 모듈로 연산
- ✓ 0으로 나누기 예외 처리

### Python 테스트

```bash
python3 test_python.py
```

**테스트 커버리지:**
- ✓ Fibonacci 수열 생성
- ✓ 소수 판별 알고리즘

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

2. **테스트 실행**
   ```bash
   # Node.js 테스트
   npm test

   # Python 테스트
   python3 test_python.py
   ```

### 새로운 실험 추가

1. 새로운 폴더를 만들어 기능/시스템 구분
2. 테스트 파일 작성 (TDD 방식)
3. 코드 구현
4. 테스트 실행으로 검증
5. Git 커밋 후 푸시하면 자동 CI/CD 실행

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
