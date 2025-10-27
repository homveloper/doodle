#!/usr/bin/env python3
"""Python unit testing verification"""

def fibonacci(n):
    """Generate fibonacci number at position n"""
    if n <= 0:
        return 0
    elif n == 1:
        return 1
    else:
        return fibonacci(n-1) + fibonacci(n-2)

def is_prime(n):
    """Check if a number is prime"""
    if n < 2:
        return False
    for i in range(2, int(n**0.5) + 1):
        if n % i == 0:
            return False
    return True

# Simple test runner
def run_tests():
    tests_passed = 0
    tests_failed = 0

    print("Running Python tests...\n")

    # Test fibonacci
    tests = [
        (fibonacci(0), 0, "fibonacci(0) should be 0"),
        (fibonacci(1), 1, "fibonacci(1) should be 1"),
        (fibonacci(5), 5, "fibonacci(5) should be 5"),
        (fibonacci(7), 13, "fibonacci(7) should be 13"),
        (is_prime(2), True, "is_prime(2) should be True"),
        (is_prime(17), True, "is_prime(17) should be True"),
        (is_prime(4), False, "is_prime(4) should be False"),
        (is_prime(1), False, "is_prime(1) should be False"),
    ]

    for actual, expected, message in tests:
        if actual == expected:
            print(f"✓ {message}")
            tests_passed += 1
        else:
            print(f"✗ {message} - Expected {expected}, got {actual}")
            tests_failed += 1

    print(f"\n{'=' * 40}")
    print(f"Total: {tests_passed + tests_failed} tests")
    print(f"Passed: {tests_passed}")
    print(f"Failed: {tests_failed}")
    print(f"{'=' * 40}")

    return tests_failed == 0

if __name__ == "__main__":
    import sys
    success = run_tests()
    sys.exit(0 if success else 1)
