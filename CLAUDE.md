# 🤖 Claude Code Web과 함께하는 개발

이 문서는 **Claude Code Web**을 활용하여 이 프로젝트를 개발하는 방법과 워크플로우를 설명합니다.

## 🌟 Claude Code Web이란?

[Claude Code Web](https://claude.com/claude-code)은 Anthropic의 최신 AI 코딩 도구로, 웹 브라우저에서 직접 코드를 작성하고 테스트할 수 있는 환경을 제공합니다. 별도의 설치 없이 브라우저만으로 전체 개발 워크플로우를 수행할 수 있습니다.

### 주요 특징

- **브라우저 기반 개발 환경**: 설치 없이 즉시 시작
- **AI 페어 프로그래밍**: Claude와 함께 코드 작성
- **멀티 언어 지원**: Node.js, Python 등 다양한 언어
- **통합 테스팅**: TDD를 바로 실행 가능
- **Git 통합**: GitHub와 직접 연동

## 🎯 이 프로젝트에서의 활용

### 1. 실험 중심 개발

Doodle 프로젝트는 새로운 아이디어를 빠르게 시도하고 검증하는 데 최적화되어 있습니다:

```
💡 아이디어 → 📝 테스트 작성 → 💻 구현 → ✅ 자동 검증
```

### 2. TDD 워크플로우

Claude Code Web을 사용한 TDD 개발 사이클:

1. **테스트 먼저 작성**
   ```javascript
   // 새로운 기능에 대한 테스트 작성
   assertEqual(newFunction(input), expectedOutput, 'description');
   ```

2. **Claude에게 구현 요청**
   ```
   "위 테스트를 통과하는 newFunction을 구현해줘"
   ```

3. **자동 실행 및 검증**
   ```bash
   npm test  # 또는 python3 test_python.py
   ```

4. **CI/CD 자동 확인**
   - Git 푸시 시 자동으로 여러 환경에서 테스트

### 3. Feature 중심 구조

프로젝트는 **feature별로** 폴더를 구분하며, 각 feature는 **가장 적합한 언어**를 자유롭게 선택합니다:

```
features/
├── calculator-js/          # JavaScript로 구현한 계산기
├── algorithms-python/      # Python으로 구현한 알고리즘
├── web-server-go/         # Go로 만든 웹 서버 (예정)
├── data-viz-rust/         # Rust로 만든 데이터 시각화 (예정)
└── ...                    # 언어에 구애받지 않는 실험들
```

**핵심 원칙:**
- 각 feature는 완전히 독립적
- 언어는 feature의 목적에 따라 선택
- 모든 feature는 자체 테스트 포함
- README에 언어와 목적 명시

## 💡 Claude Code Web 활용 팁

### 효과적인 프롬프트 작성

**좋은 예:**
```
"calculator.js에 제곱근 함수를 추가하고,
음수 입력에 대한 예외 처리를 포함한 테스트를 작성해줘.
TDD 방식으로 먼저 테스트부터 작성하자."
```

**피해야 할 예:**
```
"코드 좀 짜줘"  # 너무 모호함
```

### 단계별 개발 요청

복잡한 기능은 단계별로 나누어 요청:

1. "먼저 기본 구조만 만들어줘"
2. "이제 에러 처리를 추가하자"
3. "마지막으로 엣지 케이스 테스트를 추가해줘"

### 코드 리뷰 요청

```
"방금 작성한 코드를 리뷰해줘.
특히 성능과 에러 처리 부분을 중점적으로 봐줘."
```

## 🔄 개발 워크플로우

### 새로운 Feature 추가

```bash
# 1. 언어를 명시하여 Claude에게 요청
"새로운 문자열 처리 유틸리티를 Rust로 만들고 싶어.
features/string-utils-rust 폴더를 만들고,
테스트와 함께 구현해줘."

# 2. Claude가 자동으로:
#    - features/string-utils-rust/ 폴더 생성
#    - 필요한 설정 파일 생성 (Cargo.toml 등)
#    - 테스트 작성
#    - 구현
#    - README.md 작성

# 3. 로컬에서 테스트
cd features/string-utils-rust
cargo test

# 4. 커밋 및 푸시 (Claude가 자동 처리)
# 5. CI/CD 자동 실행
# 6. Pull Request 생성
```

### 버그 수정

```bash
# 1. 이슈 재현
"calculator.js의 divide 함수가
Infinity를 반환하는 케이스가 있어. 테스트 추가하고 수정해줘."

# 2. 테스트 작성 확인
# 3. 수정 사항 확인
# 4. 모든 테스트 통과 확인
npm test

# 5. 커밋
```

## 🚀 고급 활용

### 멀티 언어 프로젝트

이 프로젝트의 강점은 **언어에 구애받지 않는다**는 점입니다. Claude에게 자유롭게 요청하세요:

```
"HTTP 서버를 만들고 싶은데, Go와 Node.js 두 가지 버전으로 만들어줘.
각각 features/http-server-go, features/http-server-js에 만들고,
성능 비교도 할 수 있게 벤치마크 테스트도 추가해줘."
```

```
"머신러닝 예제를 Python으로 만들고 싶어.
features/ml-basics-python 폴더에 만들어줘."
```

```
"시스템 프로그래밍을 배우고 싶어서 Rust로 간단한
파일 시스템 유틸리티를 만들고 싶어."
```

**언어 선택 가이드:**
- **JavaScript/Node.js**: 웹, API, 빠른 프로토타이핑
- **Python**: 데이터, ML, 스크립팅
- **Rust**: 시스템, 성능, 메모리 안정성
- **Go**: 동시성, 서버, CLI 도구
- **TypeScript**: 타입 안정성이 중요한 프로젝트

### CI/CD 활용

GitHub Actions가 자동으로:
- ✅ 여러 Node.js 버전에서 테스트 (18, 20, 22)
- ✅ 여러 Python 버전에서 테스트 (3.9-3.12)
- ✅ 모든 테스트 결과를 요약하여 표시

### 문서 자동 생성

```
"방금 만든 함수들에 대한 API 문서를 마크다운으로 생성해줘."
```

## 📚 학습 자료

### Claude Code Web 공식 문서
- [Getting Started](https://docs.claude.com/claude-code)
- [Best Practices](https://docs.claude.com/claude-code/best-practices)
- [Examples](https://docs.claude.com/claude-code/examples)

### TDD 리소스
- [TDD with JavaScript](https://github.com/dwyl/learn-tdd)
- [TDD with Python](https://testdriven.io/test-driven-development/)

### GitHub Actions
- [Workflow Syntax](https://docs.github.com/actions/using-workflows/workflow-syntax-for-github-actions)
- [Best Practices](https://docs.github.com/actions/learn-github-actions/usage-limits-billing-and-administration)

## 💬 프롬프트 예제 모음

### 기능 개발
```
"회문(palindrome) 체크 함수를 만들고 싶어.
1. 먼저 테스트 케이스를 작성해줘
2. 그 다음 함수를 구현하고
3. 마지막으로 엣지 케이스를 처리해줘"
```

### 리팩토링
```
"calculator.js의 코드를 더 함수형 프로그래밍 스타일로
리팩토링해줘. 기존 테스트는 모두 통과해야 해."
```

### 성능 최적화
```
"fibonacci 함수가 너무 느려. 메모이제이션을 적용해서
성능을 개선하고, 벤치마크 테스트도 추가해줘."
```

### 통합 요청
```
"새로운 feature를 C++로 만들고 싶어.
features/data-structures-cpp 폴더에 스택과 큐를 구현하고,
Google Test로 테스트 작성해줘.
CI/CD 파이프라인에도 자동으로 포함되게 설정해줘."
```

### 언어별 Feature 요청
```
"features/ 폴더에 다음 feature들을 만들어줘:
1. features/web-scraper-python - BeautifulSoup 사용
2. features/rest-api-typescript - Express + TypeScript
3. features/cli-tool-go - 간단한 파일 검색 도구"
```

## 🎓 배운 점과 팁

### 효과적인 협업
- **명확한 요구사항**: Claude에게 구체적으로 설명할수록 더 좋은 결과
- **단계별 접근**: 큰 기능은 작은 단계로 나누기
- **테스트 우선**: TDD를 따르면 더 견고한 코드

### 시간 절약
- **자동화**: CI/CD로 반복 작업 자동화
- **빠른 반복**: 아이디어를 즉시 코드로 변환
- **학습 가속**: 다양한 언어와 패턴을 빠르게 실험

### 코드 품질
- **자동 테스트**: 모든 변경사항이 자동으로 검증
- **다중 환경**: 여러 버전에서 호환성 확인
- **문서화**: 코드와 함께 문서도 함께 관리

## 🤝 기여하기

새로운 실험 아이디어가 있다면:
1. 새로운 폴더 생성
2. README 작성
3. 테스트 작성
4. 구현
5. Pull Request

Claude Code Web과 함께라면 이 모든 과정이 훨씬 빠르고 즐겁습니다!

---

**Happy Coding with Claude! 🎉**
