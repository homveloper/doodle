# Doodle Project Conventions

This document outlines the conventions and patterns observed in the Doodle project to guide feature creation.

## Feature Naming Conventions

### Folder Names

Features follow a consistent naming pattern: `{feature-name}-{language-suffix}`

**Examples from existing features:**
- `calculator-js` - JavaScript calculator
- `algorithms-python` - Python algorithms collection
- `ascii-art-go` - Go ASCII art generator
- `hellojava` - Java/Kotlin hello world (Java often omits suffix)

**Pattern Rules:**
- Use lowercase kebab-case
- Add language suffix for clarity (optional for Java)
- Keep names descriptive but concise
- Suffix patterns: `-js`, `-python`, `-go`, `-rust`, `-cpp`, etc.

### When Feature Already Exists

If a feature folder with the same name already exists:
1. Inform the user
2. Continue working within the existing folder
3. Add to or modify existing files
4. Don't recreate or duplicate structure

## Language-Specific Patterns

### JavaScript/Node.js

**Typical Structure:**
```
feature-name-js/
├── {feature}.js       # Implementation
├── test.js           # Tests with custom runner
├── package.json      # Minimal config
└── README.md         # Documentation
```

**Key Characteristics:**
- Minimal `package.json` (name, version, test script)
- Custom test runner (no external frameworks preferred)
- Simple assertion-based tests
- CommonJS or ES modules depending on preference
- Test command: `npm test`

**Example package.json:**
```json
{
  "name": "feature-name",
  "version": "1.0.0",
  "scripts": {
    "test": "node test.js"
  }
}
```

### Python

**Typical Structure:**
```
feature-name-python/
├── test_{feature}.py  # Combined implementation + tests
└── README.md         # Documentation
```

**Key Characteristics:**
- Tests and implementation often in same file
- No external dependencies (use stdlib)
- Simple assert-based testing
- Python 3.9+ compatibility required
- Test command: `python3 test_{feature}.py`

**Test Pattern:**
```python
# Implementation
def my_function():
    pass

# Tests
def test_my_function():
    assert my_function() == expected
    print("✓ Test passed")

if __name__ == "__main__":
    run_tests()
```

### Go

**Typical Structure:**
```
feature-name-go/
├── go.mod
├── {package}/
│   ├── {package}.go
│   └── {package}_test.go
├── examples/          # Optional
└── README.md
```

**Key Characteristics:**
- Module path: `github.com/homveloper/doodle/features/{feature-name}`
- Go version: 1.24+
- Package-based structure
- Standard Go test framework
- Test command: `go test -v`

**go.mod Example:**
```go
module github.com/homveloper/doodle/features/feature-name-go

go 1.24
```

### Java/Gradle

**Typical Structure:**
```
feature-name/
├── settings.gradle.kts
├── gradle/            # Wrapper
├── gradlew*
├── app/
│   ├── build.gradle.kts
│   └── src/
│       ├── main/java/{package}/
│       │   └── App.java
│       └── test/java/{package}/
│           └── AppTest.java
└── README.md
```

**Key Characteristics:**
- Multi-module Gradle project
- Java 21+ toolchain
- JUnit 5 for testing
- Package names match folder (without hyphens)
- Test command: `./gradlew test`

## README.md Structure

All features must have a comprehensive README with these sections:

### 1. Title & Metadata
```markdown
# Feature Name (Language)

**Language**: JavaScript/Python/Go/Java
**Purpose**: Brief description
**Status**: 🚧 In Progress / ✅ Complete
```

### 2. Overview
Brief description of what the feature does and why it exists.

### 3. Features List
Bullet points or checklist of implemented functionality.

### 4. Quick Start
Installation and setup instructions:
- Prerequisites (language version, tools)
- Installation commands
- How to run tests
- Basic usage example

### 5. Usage Examples
Code examples showing how to use the feature.

### 6. Test Coverage
Description of tests and what they verify.

### 7. Project Structure
ASCII tree showing file organization.

### 8. Future Plans
Checklist of upcoming features or improvements.

### 9. Development Notes (Optional)
Performance considerations, design decisions, limitations.

## TDD Approach

All features should follow Test-Driven Development:

1. **Write tests first** - Define expected behavior
2. **Implement functionality** - Make tests pass
3. **Refactor** - Improve code while keeping tests green
4. **Document** - Update README with examples

## CI/CD Integration

Features are automatically tested via GitHub Actions:

**Current Support:**
- Node.js: Versions 18.x, 20.x, 22.x
- Python: Versions 3.9, 3.10, 3.11, 3.12

**To Add Support for Go/Java:**
The CI workflow (`.github/workflows/test.yml`) can be extended to include:
- Go testing
- Gradle/Java testing

## Best Practices

### Dependencies
- **Prefer zero dependencies** for simplicity
- Use standard library when possible
- External libraries only when necessary
- Document why external dependencies are needed

### Code Organization
- Keep files focused and single-purpose
- Use language-specific naming conventions
- Follow established patterns from existing features
- Maintain consistency within each language

### Testing
- Write clear, descriptive test names
- Test both success and failure cases
- Include edge cases
- Keep tests simple and readable

### Documentation
- README must be comprehensive
- Include code examples
- Document prerequisites clearly
- Explain design decisions when non-obvious

## Existing Features Reference

| Feature | Language | Characteristics |
|---------|----------|-----------------|
| `calculator-js` | JavaScript | Simple arithmetic, custom test runner, zero deps |
| `algorithms-python` | Python | Data structures, combined impl+test file, stdlib only |
| `ascii-art-go` | Go | Package structure, examples included, standard testing |
| `hellojava` | Java | Gradle multi-module, JUnit 5, demonstrates Java ecosystem |

## Creating New Features

When asked to create a new feature:

1. **Determine language** - From explicit request or infer from context
2. **Check existing folder** - Look in `features/` for existing work
3. **Create appropriate structure** - Follow language-specific patterns
4. **Start with tests** - TDD approach
5. **Generate README** - Complete documentation
6. **Consider CI/CD** - Whether new workflow steps needed

**Key Principle**: Adapt to the feature's needs rather than forcing a rigid structure. Each feature is an experiment and should evolve naturally.
