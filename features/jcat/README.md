# JCat - File Information CLI Tool

Java로 구현한 파일 정보 표시 CLI 도구입니다. 모든 에러 핸들링은 **error-as-value** 패턴(`Result<T, E>`)으로 처리됩니다.

## Features

- **파일 읽기**: 파일 내용을 터미널에 표시
- **파일 정보**: 파일 메타데이터 (크기, 권한, 수정 시간 등) 표시
- **Type-safe 에러 핸들링**: Exception 대신 Result 타입 사용
- **명확한 에러 메시지**: 파일을 찾을 수 없음, 읽을 수 없음 등 구체적 에러

## Installation

### Build

```bash
./gradlew jar
```

### Run

```bash
# 직접 실행
java -jar build/libs/jcat-1.0.0.jar <command> <file>

# 또는 wrapper 스크립트 사용
./jcat <command> <file>
```

### Optional: Add to PATH

```bash
# ~/.bashrc 또는 ~/.zshrc에 추가
export PATH="$PATH:/path/to/doodle/features/jcat"
```

## Usage

### Commands

#### `read` - 파일 내용 읽기

```bash
./jcat read ./sample.txt
# 출력: 파일의 전체 내용
```

#### `info` - 파일 정보 표시

```bash
./jcat info ./sample.txt
# 출력:
# Path: /absolute/path/to/sample.txt
# Size: 1.23 KB (1259 bytes)
# Type: Regular File
# Permissions: rw-
# Last Modified: 2025-11-05 18:28:27
```

#### `help` - 도움말

```bash
./jcat help
./jcat -h
./jcat --help
```

## Architecture

### Error-as-Value Pattern

모든 파일 작업은 `Result<T, E>` 타입을 반환하여 에러를 명시적으로 처리합니다:

```java
// ✅ Good: Error-as-value
Result<String, FileError> result = FileOperations.readFile("test.txt");
result.match(
    content -> System.out.println(content),
    error -> System.err.println(error.message())
);

// ❌ Bad: Exception-based (사용 안 함)
try {
    String content = Files.readString(Path.of("test.txt"));
} catch (IOException e) {
    // ...
}
```

### Core Types

#### `Result<T, E>`

```java
sealed interface Result<T, E> permits Ok, Err {
    record Ok<T, E>(T value) implements Result<T, E> {}
    record Err<T, E>(E error) implements Result<T, E> {}

    // Functional operations
    <U> Result<U, E> map(Function<T, U> mapper);
    <U> Result<U, E> andThen(Function<T, Result<U, E>> mapper);
    <U> U match(Function<T, U> okMapper, Function<E, U> errMapper);
}
```

#### `FileError`

```java
sealed interface FileError {
    record NotFound(String path) implements FileError {}
    record NotReadable(String path) implements FileError {}
    record IoError(String path, String message) implements FileError {}
    record InvalidPath(String path) implements FileError {}
}
```

#### `FileInfo`

파일 메타데이터를 담는 record:

```java
record FileInfo(
    String path,
    long size,
    boolean isDirectory,
    boolean isRegularFile,
    boolean isReadable,
    boolean isWritable,
    boolean isExecutable,
    FileTime lastModified
)
```

### Project Structure

```
jcat/
├── src/
│   ├── main/java/com/doodle/jcat/
│   │   ├── Result.java           # Generic Result<T, E> type
│   │   ├── FileError.java         # File operation errors
│   │   ├── FileInfo.java          # File metadata record
│   │   ├── FileOperations.java    # Core file operations
│   │   ├── Command.java           # CLI command types
│   │   ├── CommandParser.java     # Argument parsing
│   │   └── JCat.java             # Main application
│   └── test/java/com/doodle/jcat/
│       ├── FileOperationsTest.java    # 11 tests
│       ├── CommandParserTest.java     # 7 tests
│       └── JCatTest.java             # 5 tests
├── build.gradle
├── settings.gradle
├── jcat                          # Launcher script
└── README.md
```

## Development

### Run Tests

```bash
./gradlew test
```

모든 테스트는 TDD 방식으로 작성되었으며 **23개 테스트**가 포함되어 있습니다.

### Test Coverage

- **FileOperations**: 파일 읽기, 정보 조회, 에러 처리
- **Result Pattern**: map, andThen, match 체이닝
- **CommandParser**: 모든 명령어 파싱 케이스
- **JCat CLI**: 성공/실패 시나리오

### Examples from Tests

```java
// Result chaining with map
Result<Integer, FileError> lengthResult =
    FileOperations.readFile("test.txt")
        .map(String::length);

// Result chaining with andThen
Result<FileInfo, FileError> result =
    FileOperations.readFile("path.txt")
        .andThen(path -> FileOperations.getFileInfo(path.trim()));

// Pattern matching
String message = FileOperations.readFile("test.txt")
    .match(
        content -> "Success: " + content.length() + " chars",
        error -> "Error: " + error.message()
    );
```

## Requirements

- Java 21+
- Gradle 9.0+ (wrapper included)

## Why Error-as-Value?

### 장점

1. **타입 안정성**: 컴파일 타임에 에러 처리 강제
2. **명시적**: 함수 시그니처에서 실패 가능성 명확히 표현
3. **조합 가능**: `map`, `andThen`으로 함수형 체이닝
4. **제어 흐름 단순화**: try-catch 중첩 제거
5. **에러 타입 문서화**: Sealed interface로 모든 에러 케이스 명시

### vs Exception

| Error-as-Value | Exception |
|---------------|-----------|
| `Result<T, FileError>` | `throws IOException` |
| 명시적 에러 타입 | 런타임 에러 가능 |
| 함수형 체이닝 | try-catch 중첩 |
| 컴파일 타임 체크 | 런타임 발견 |

## License

MIT

## Related

이 프로젝트는 Doodle의 일부로, **error-as-value** 패턴을 Java에서 실용적으로 적용하는 예제입니다.

`.claude/skills/java-error-as-value/` 스킬과 함께 사용하면 다른 프로젝트에서도 동일한 패턴을 쉽게 적용할 수 있습니다.
