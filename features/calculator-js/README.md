# Calculator (JavaScript)

**언어**: JavaScript (Node.js)
**목적**: TDD를 활용한 기본적인 수학 연산 함수 구현

## 개요

외부 의존성 없이 순수 JavaScript로 작성된 계산기 라이브러리입니다. 자체 테스트 러너를 사용하여 TDD 방식으로 개발되었습니다.

## 기능

### 구현된 연산
- `add(a, b)` - 덧셈
- `subtract(a, b)` - 뺄셈
- `multiply(a, b)` - 곱셈
- `divide(a, b)` - 나눗셈 (0으로 나누기 예외 처리)
- `power(a, b)` - 거듭제곱
- `modulo(a, b)` - 나머지 연산

## 사용 방법

### 설치
```bash
cd features/calculator-js
```

### 테스트 실행
```bash
npm test
```

### 코드 사용 예제
```javascript
const { add, subtract, multiply, divide, power, modulo } = require('./calculator.js');

console.log(add(5, 3));        // 8
console.log(subtract(10, 4));  // 6
console.log(multiply(3, 7));   // 21
console.log(divide(15, 3));    // 5
console.log(power(2, 3));      // 8
console.log(modulo(10, 3));    // 1
```

## 테스트 커버리지

총 15개의 테스트 케이스:
- ✓ 덧셈: 양수, 음수, 0 케이스
- ✓ 뺄셈: 양수, 음수 케이스
- ✓ 곱셈: 양수, 음수 케이스
- ✓ 나눗셈: 정수, 소수, 0으로 나누기 예외
- ✓ 거듭제곱: 일반, 0제곱 케이스
- ✓ 모듈로: 다양한 나머지 케이스

## 구조

```
calculator-js/
├── calculator.js    # 계산기 로직
├── test.js         # 테스트 케이스
├── package.json    # 프로젝트 설정
└── README.md       # 이 파일
```

## 향후 계획

- [ ] 삼각함수 추가 (sin, cos, tan)
- [ ] 로그 함수 추가
- [ ] 복잡한 수 연산 지원
- [ ] 더 많은 엣지 케이스 테스트
