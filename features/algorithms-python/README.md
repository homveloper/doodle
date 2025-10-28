# Algorithms (Python)

**언어**: Python 3.9+
**목적**: 기본 알고리즘 구현 및 검증

## 개요

Python으로 작성된 기본 알고리즘 모음입니다. 각 알고리즘은 단위 테스트와 함께 제공되어 정확성을 보장합니다.

## 구현된 알고리즘

### 1. Fibonacci 수열
```python
def fibonacci(n):
    """n번째 피보나치 수를 반환"""
```

**특징:**
- 재귀적 구현
- 시간 복잡도: O(2^n)
- 추후 메모이제이션으로 최적화 예정

**사용 예제:**
```python
print(fibonacci(0))   # 0
print(fibonacci(1))   # 1
print(fibonacci(5))   # 5
print(fibonacci(10))  # 55
```

### 2. 소수 판별
```python
def is_prime(n):
    """n이 소수인지 확인"""
```

**특징:**
- 제곱근까지만 확인하는 최적화
- 시간 복잡도: O(√n)
- 2 미만의 수는 소수가 아님

**사용 예제:**
```python
print(is_prime(2))    # True
print(is_prime(17))   # True
print(is_prime(4))    # False
print(is_prime(1))    # False
```

## 사용 방법

### 설치
```bash
cd features/algorithms-python
```

Python 3.9 이상이 필요합니다.

### 테스트 실행
```bash
python3 test_python.py
```

### 직접 사용
```python
from test_python import fibonacci, is_prime

# Fibonacci 사용
result = fibonacci(7)
print(f"7번째 피보나치 수: {result}")

# 소수 판별 사용
numbers = [2, 3, 4, 5, 6, 7, 8, 9, 10]
primes = [n for n in numbers if is_prime(n)]
print(f"소수들: {primes}")
```

## 테스트 커버리지

총 8개의 테스트 케이스:
- ✓ Fibonacci: 0, 1, 5, 7 케이스
- ✓ 소수 판별: 2, 17 (소수), 4, 1 (비소수)

## 구조

```
algorithms-python/
├── test_python.py    # 알고리즘 구현 및 테스트
└── README.md        # 이 파일
```

## 향후 계획

- [ ] Fibonacci 메모이제이션 최적화
- [ ] 정렬 알고리즘 추가 (퀵소트, 머지소트)
- [ ] 탐색 알고리즘 추가 (이진 탐색)
- [ ] 그래프 알고리즘 추가 (DFS, BFS)
- [ ] 동적 프로그래밍 예제 추가
- [ ] 성능 벤치마크 추가

## 성능 참고사항

**Fibonacci 함수:**
- 현재 재귀 방식이라 n > 35에서 느려질 수 있음
- 큰 수를 계산할 경우 메모이제이션 버전 사용 권장 (추후 추가 예정)

**is_prime 함수:**
- 큰 수에서도 효율적으로 작동
- n = 1,000,000 정도까지 빠르게 판별 가능
