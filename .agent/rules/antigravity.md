---
trigger: always_on
---

# Doodle Project

**실험적 feature 기반 멀티 언어 프로젝트**

## 구조

```
features/
├── calculator-js/       # JavaScript 계산기
├── algorithms-python/   # Python 알고리즘
├── ascii-art-go/       # Go ASCII 아트
├── hellojava/          # Java 실험
└── ...                 # 언어 자유 선택
```

## 핵심 원칙

- **Feature 독립성**: 각 feature는 완전히 독립적
- **언어 자유**: 목적에 맞는 언어 선택
- **TDD 우선**: 테스트 먼저, 구현 나중
- **자체 문서화**: 각 feature에 README 필수

## Feature 명명 규칙

- `{feature-name}-{language}`: `calculator-js`, `algorithms-python`
- Java는 언어 접미사 생략 가능: `hellojava`

## 새 Feature 추가

```
"Python으로 웹스크래퍼 doodle 만들어줘"
→ features/web-scraper-python/ 생성
→ TDD로 구현
→ README 작성
```

## CI/CD

- Node.js: 18.x, 20.x, 22.x
- Python: 3.9-3.12
- 자동 테스트 및 검증

## 언어별 구조

**JavaScript**: `package.json`, `test.js`, `{feature}.js`
**Python**: `test_{feature}.py` (구현+테스트 통합)
**Go**: `go.mod`, `{pkg}/{pkg}.go`, `{pkg}_test.go`
**Java**: Gradle 멀티모듈, `src/main/java`, `src/test/java`

---

자세한 내용은 `.claude/skills/doodle/` 참고
