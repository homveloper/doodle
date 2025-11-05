# Language-Specific Examples

This document provides concrete examples of how each language is typically structured in Doodle features.

## JavaScript Example

### Simple Feature: String Utilities

**Folder**: `string-utils-js/`

**package.json**:
```json
{
  "name": "string-utils",
  "version": "1.0.0",
  "description": "String manipulation utilities",
  "main": "string_utils.js",
  "scripts": {
    "test": "node test.js"
  },
  "keywords": ["doodle", "strings", "utilities"],
  "license": "MIT"
}
```

**string_utils.js**:
```javascript
/**
 * String manipulation utilities
 */

function reverse(str) {
    return str.split('').reverse().join('');
}

function isPalindrome(str) {
    const cleaned = str.toLowerCase().replace(/[^a-z0-9]/g, '');
    return cleaned === reverse(cleaned);
}

function capitalize(str) {
    return str.charAt(0).toUpperCase() + str.slice(1);
}

module.exports = {
    reverse,
    isPalindrome,
    capitalize
};
```

**test.js**:
```javascript
const { reverse, isPalindrome, capitalize } = require('./string_utils');

function assertEqual(actual, expected, testName) {
    if (JSON.stringify(actual) === JSON.stringify(expected)) {
        console.log(`✓ ${testName}`);
        return true;
    } else {
        console.error(`✗ ${testName}`);
        console.error(`  Expected: ${expected}, Got: ${actual}`);
        return false;
    }
}

console.log('Running string utilities tests...\n');

let passed = 0;
let failed = 0;

// Test reverse
if (assertEqual(reverse('hello'), 'olleh', 'reverse() with simple string')) passed++;
else failed++;

if (assertEqual(reverse(''), '', 'reverse() with empty string')) passed++;
else failed++;

// Test isPalindrome
if (assertEqual(isPalindrome('racecar'), true, 'isPalindrome() with valid palindrome')) passed++;
else failed++;

if (assertEqual(isPalindrome('hello'), false, 'isPalindrome() with non-palindrome')) passed++;
else failed++;

if (assertEqual(isPalindrome('A man a plan a canal Panama'), true, 'isPalindrome() ignores spaces and case')) passed++;
else failed++;

// Test capitalize
if (assertEqual(capitalize('hello'), 'Hello', 'capitalize() simple word')) passed++;
else failed++;

console.log(`\n${passed} passed, ${failed} failed`);
process.exit(failed > 0 ? 1 : 0);
```

## Python Example

### Simple Feature: Math Utilities

**Folder**: `math-utils-python/`

**test_math_utils.py**:
```python
"""
Math Utilities - Common mathematical operations

This file contains both implementation and tests following TDD approach.
"""

import math

# ============================================================================
# Implementation
# ============================================================================

def factorial(n):
    """Calculate factorial of n"""
    if n < 0:
        raise ValueError("Factorial not defined for negative numbers")
    if n == 0 or n == 1:
        return 1
    return n * factorial(n - 1)

def is_prime(n):
    """Check if number is prime"""
    if n < 2:
        return False
    if n == 2:
        return True
    if n % 2 == 0:
        return False
    for i in range(3, int(math.sqrt(n)) + 1, 2):
        if n % i == 0:
            return False
    return True

def fibonacci(n):
    """Generate first n Fibonacci numbers"""
    if n <= 0:
        return []
    if n == 1:
        return [0]
    fib = [0, 1]
    for i in range(2, n):
        fib.append(fib[i-1] + fib[i-2])
    return fib

# ============================================================================
# Tests
# ============================================================================

def test_factorial():
    """Test factorial function"""
    assert factorial(0) == 1, "Factorial of 0 should be 1"
    assert factorial(1) == 1, "Factorial of 1 should be 1"
    assert factorial(5) == 120, "Factorial of 5 should be 120"
    assert factorial(10) == 3628800, "Factorial of 10 should be 3628800"
    print("✓ factorial() tests passed")

def test_is_prime():
    """Test prime number checker"""
    assert is_prime(2) == True, "2 is prime"
    assert is_prime(3) == True, "3 is prime"
    assert is_prime(4) == False, "4 is not prime"
    assert is_prime(17) == True, "17 is prime"
    assert is_prime(100) == False, "100 is not prime"
    assert is_prime(1) == False, "1 is not prime"
    print("✓ is_prime() tests passed")

def test_fibonacci():
    """Test Fibonacci generator"""
    assert fibonacci(0) == [], "Fibonacci of 0 should be empty"
    assert fibonacci(1) == [0], "Fibonacci of 1 should be [0]"
    assert fibonacci(5) == [0, 1, 1, 2, 3], "First 5 Fibonacci numbers"
    assert fibonacci(8) == [0, 1, 1, 2, 3, 5, 8, 13], "First 8 Fibonacci numbers"
    print("✓ fibonacci() tests passed")

def run_tests():
    """Run all tests"""
    print("Running math utilities tests...\n")

    tests = [
        test_factorial,
        test_is_prime,
        test_fibonacci
    ]

    passed = 0
    failed = 0

    for test in tests:
        try:
            test()
            passed += 1
        except AssertionError as e:
            print(f"✗ {test.__name__} failed: {e}")
            failed += 1
        except Exception as e:
            print(f"✗ {test.__name__} error: {e}")
            failed += 1

    print(f"\n{passed} passed, {failed} failed")
    return failed == 0

if __name__ == "__main__":
    import sys
    sys.exit(0 if run_tests() else 1)
```

## Go Example

### Simple Feature: File Utilities

**Folder**: `file-utils-go/`

**go.mod**:
```go
module github.com/homveloper/doodle/features/file-utils-go

go 1.24
```

**fileutils/fileutils.go**:
```go
// Package fileutils provides utilities for file operations
package fileutils

import (
	"bufio"
	"os"
	"strings"
)

// CountLines counts the number of lines in a file
func CountLines(filepath string) (int, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		count++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return count, nil
}

// CountWords counts the number of words in a file
func CountWords(filepath string) (int, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return count, nil
}

// ReadLines reads all lines from a file
func ReadLines(filepath string) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

// SearchText searches for a substring in a file
func SearchText(filepath, search string) ([]int, error) {
	lines, err := ReadLines(filepath)
	if err != nil {
		return nil, err
	}

	var lineNumbers []int
	for i, line := range lines {
		if strings.Contains(line, search) {
			lineNumbers = append(lineNumbers, i+1)
		}
	}

	return lineNumbers, nil
}
```

**fileutils/fileutils_test.go**:
```go
package fileutils

import (
	"os"
	"testing"
)

func TestCountLines(t *testing.T) {
	// Create test file
	content := "line 1\nline 2\nline 3\n"
	tmpfile, err := os.CreateTemp("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write([]byte(content)); err != nil {
		t.Fatal(err)
	}
	tmpfile.Close()

	// Test
	count, err := CountLines(tmpfile.Name())
	if err != nil {
		t.Fatalf("CountLines failed: %v", err)
	}

	expected := 3
	if count != expected {
		t.Errorf("Expected %d lines, got %d", expected, count)
	}
	t.Log("✓ CountLines test passed")
}

func TestCountWords(t *testing.T) {
	content := "hello world this is a test"
	tmpfile, err := os.CreateTemp("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write([]byte(content)); err != nil {
		t.Fatal(err)
	}
	tmpfile.Close()

	count, err := CountWords(tmpfile.Name())
	if err != nil {
		t.Fatalf("CountWords failed: %v", err)
	}

	expected := 6
	if count != expected {
		t.Errorf("Expected %d words, got %d", expected, count)
	}
	t.Log("✓ CountWords test passed")
}

func TestReadLines(t *testing.T) {
	content := "line 1\nline 2\nline 3"
	tmpfile, err := os.CreateTemp("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write([]byte(content)); err != nil {
		t.Fatal(err)
	}
	tmpfile.Close()

	lines, err := ReadLines(tmpfile.Name())
	if err != nil {
		t.Fatalf("ReadLines failed: %v", err)
	}

	expected := 3
	if len(lines) != expected {
		t.Errorf("Expected %d lines, got %d", expected, len(lines))
	}
	t.Log("✓ ReadLines test passed")
}

func TestSearchText(t *testing.T) {
	content := "hello world\ntest line\nhello again\nfinal line"
	tmpfile, err := os.CreateTemp("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write([]byte(content)); err != nil {
		t.Fatal(err)
	}
	tmpfile.Close()

	lineNumbers, err := SearchText(tmpfile.Name(), "hello")
	if err != nil {
		t.Fatalf("SearchText failed: %v", err)
	}

	expected := []int{1, 3}
	if len(lineNumbers) != len(expected) {
		t.Errorf("Expected %v, got %v", expected, lineNumbers)
	}
	t.Log("✓ SearchText test passed")
}
```

## Java Example

### Simple Feature: Collections Utilities

**Folder**: `collections-utils/`

**settings.gradle.kts**:
```kotlin
rootProject.name = "collections-utils"
include("app")
```

**app/build.gradle.kts**:
```kotlin
plugins {
    application
    java
}

repositories {
    mavenCentral()
}

dependencies {
    testImplementation("org.junit.jupiter:junit-jupiter:5.10.0")
    testRuntimeOnly("org.junit.platform:junit-platform-launcher")
    implementation("com.google.guava:guava:32.1.1-jre")
}

java {
    toolchain {
        languageVersion = JavaLanguageVersion.of(21)
    }
}

application {
    mainClass = "collectionsutils.App"
}

tasks.named<Test>("test") {
    useJUnitPlatform()
}
```

**app/src/main/java/collectionsutils/CollectionUtils.java**:
```java
package collectionsutils;

import java.util.*;
import java.util.stream.Collectors;

/**
 * Utility functions for collection operations
 */
public class CollectionUtils {

    /**
     * Remove duplicates from a list while preserving order
     */
    public static <T> List<T> removeDuplicates(List<T> list) {
        return new ArrayList<>(new LinkedHashSet<>(list));
    }

    /**
     * Partition a list into chunks of specified size
     */
    public static <T> List<List<T>> partition(List<T> list, int size) {
        if (size <= 0) {
            throw new IllegalArgumentException("Partition size must be positive");
        }

        List<List<T>> partitions = new ArrayList<>();
        for (int i = 0; i < list.size(); i += size) {
            partitions.add(list.subList(i, Math.min(i + size, list.size())));
        }
        return partitions;
    }

    /**
     * Find the most frequent element in a list
     */
    public static <T> Optional<T> mostFrequent(List<T> list) {
        if (list.isEmpty()) {
            return Optional.empty();
        }

        Map<T, Long> frequencies = list.stream()
            .collect(Collectors.groupingBy(e -> e, Collectors.counting()));

        return frequencies.entrySet().stream()
            .max(Map.Entry.comparingByValue())
            .map(Map.Entry::getKey);
    }

    /**
     * Flatten a list of lists into a single list
     */
    public static <T> List<T> flatten(List<List<T>> lists) {
        return lists.stream()
            .flatMap(List::stream)
            .collect(Collectors.toList());
    }

    /**
     * Zip two lists together into pairs
     */
    public static <A, B> List<Pair<A, B>> zip(List<A> listA, List<B> listB) {
        int size = Math.min(listA.size(), listB.size());
        List<Pair<A, B>> result = new ArrayList<>();

        for (int i = 0; i < size; i++) {
            result.add(new Pair<>(listA.get(i), listB.get(i)));
        }

        return result;
    }

    /**
     * Simple pair record
     */
    public record Pair<A, B>(A first, B second) {}
}
```

**app/src/test/java/collectionsutils/CollectionUtilsTest.java**:
```java
package collectionsutils;

import org.junit.jupiter.api.Test;
import java.util.*;
import static org.junit.jupiter.api.Assertions.*;

class CollectionUtilsTest {

    @Test
    void testRemoveDuplicates() {
        List<Integer> input = List.of(1, 2, 2, 3, 4, 3, 5);
        List<Integer> result = CollectionUtils.removeDuplicates(input);
        List<Integer> expected = List.of(1, 2, 3, 4, 5);

        assertEquals(expected, result);
        System.out.println("✓ removeDuplicates test passed");
    }

    @Test
    void testPartition() {
        List<Integer> input = List.of(1, 2, 3, 4, 5, 6, 7);
        List<List<Integer>> result = CollectionUtils.partition(input, 3);

        assertEquals(3, result.size());
        assertEquals(List.of(1, 2, 3), result.get(0));
        assertEquals(List.of(4, 5, 6), result.get(1));
        assertEquals(List.of(7), result.get(2));
        System.out.println("✓ partition test passed");
    }

    @Test
    void testMostFrequent() {
        List<String> input = List.of("apple", "banana", "apple", "cherry", "apple", "banana");
        Optional<String> result = CollectionUtils.mostFrequent(input);

        assertTrue(result.isPresent());
        assertEquals("apple", result.get());
        System.out.println("✓ mostFrequent test passed");
    }

    @Test
    void testFlatten() {
        List<List<Integer>> input = List.of(
            List.of(1, 2),
            List.of(3, 4, 5),
            List.of(6)
        );
        List<Integer> result = CollectionUtils.flatten(input);
        List<Integer> expected = List.of(1, 2, 3, 4, 5, 6);

        assertEquals(expected, result);
        System.out.println("✓ flatten test passed");
    }

    @Test
    void testZip() {
        List<String> names = List.of("Alice", "Bob", "Charlie");
        List<Integer> ages = List.of(25, 30, 35);

        List<CollectionUtils.Pair<String, Integer>> result = CollectionUtils.zip(names, ages);

        assertEquals(3, result.size());
        assertEquals("Alice", result.get(0).first());
        assertEquals(25, result.get(0).second());
        System.out.println("✓ zip test passed");
    }
}
```

## Common Patterns Across Languages

### Test Organization
- **JavaScript**: Custom test runner with `assertEqual` helper
- **Python**: `assert` statements with descriptive messages
- **Go**: Standard `testing` package with `t.Error/Fatal`
- **Java**: JUnit 5 with `assertEquals`, `assertTrue`, etc.

### Error Handling
- **JavaScript**: Try-catch, return null/undefined, or throw
- **Python**: Exceptions with specific error types
- **Go**: Return `(value, error)` tuples
- **Java**: Exceptions or Optional return types

### Code Organization
- **JavaScript**: Single file or multiple files with `module.exports`
- **Python**: Single file for simple features, packages for complex ones
- **Go**: Package-based, files in package directory
- **Java**: Package structure matching folder hierarchy

These examples serve as references, not rigid templates. Adapt structure based on feature requirements.
