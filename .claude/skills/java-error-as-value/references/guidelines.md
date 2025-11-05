# Error-as-Value Pattern Guidelines

## When to Use This Pattern

Use the Result<T, E> pattern when:

1. **Errors are expected and recoverable** - The operation might fail under normal conditions (e.g., parsing user input, network requests, file I/O)
2. **Caller should handle the error** - The calling code needs to decide how to handle failures
3. **Type-safe error handling is desired** - You want compile-time guarantees that errors are handled
4. **Avoiding exceptions improves clarity** - The control flow is clearer without try-catch blocks
5. **Composing fallible operations** - Chaining multiple operations that can fail

## When NOT to Use This Pattern

Avoid Result<T, E> and use exceptions when:

1. **Programming errors** - Bugs that should never happen (e.g., null pointer, array out of bounds)
2. **Unrecoverable errors** - Fatal errors that should crash the program
3. **Error handling is the same everywhere** - If all callers handle errors identically, exceptions might be clearer
4. **Legacy code integration** - When working with APIs that expect exceptions

## Designing Error Types

### Option 1: String Errors (Simple)

For quick prototyping or simple applications:

```java
Result<Integer, String> parseAge(String input) {
    try {
        int age = Integer.parseInt(input);
        if (age < 0 || age > 150) {
            return Result.err("Age must be between 0 and 150");
        }
        return Result.ok(age);
    } catch (NumberFormatException e) {
        return Result.err("Invalid number format: " + input);
    }
}
```

**Pros**: Quick to implement, easy to read
**Cons**: No type safety, hard to handle different error types programmatically

### Option 2: Enum Errors (Recommended)

For most production code:

```java
enum ParseError {
    INVALID_FORMAT,
    OUT_OF_RANGE,
    EMPTY_INPUT
}

Result<Integer, ParseError> parseAge(String input) {
    if (input == null || input.isEmpty()) {
        return Result.err(ParseError.EMPTY_INPUT);
    }
    try {
        int age = Integer.parseInt(input);
        if (age < 0 || age > 150) {
            return Result.err(ParseError.OUT_OF_RANGE);
        }
        return Result.ok(age);
    } catch (NumberFormatException e) {
        return Result.err(ParseError.INVALID_FORMAT);
    }
}
```

**Pros**: Type-safe, exhaustive matching, easy to extend
**Cons**: Less flexible for error messages

### Option 3: Sealed Interface Errors (Advanced)

For complex error hierarchies with data:

```java
sealed interface ParseError {
    record InvalidFormat(String input, String reason) implements ParseError {}
    record OutOfRange(int value, int min, int max) implements ParseError {}
    record EmptyInput() implements ParseError {}
}

Result<Integer, ParseError> parseAge(String input) {
    if (input == null || input.isEmpty()) {
        return Result.err(new ParseError.EmptyInput());
    }
    try {
        int age = Integer.parseInt(input);
        if (age < 0 || age > 150) {
            return Result.err(new ParseError.OutOfRange(age, 0, 150));
        }
        return Result.ok(age);
    } catch (NumberFormatException e) {
        return Result.err(new ParseError.InvalidFormat(input, e.getMessage()));
    }
}
```

**Pros**: Type-safe, rich error information, exhaustive matching
**Cons**: More verbose, requires Java 17+ for sealed interfaces

## Common Patterns

### 1. Wrapping Exception-Throwing Code

```java
// Before: Exception-based
public int readFileSize(String path) throws IOException {
    return (int) Files.size(Paths.get(path));
}

// After: Result-based
public Result<Integer, String> readFileSize(String path) {
    return Result.of(
        () -> (int) Files.size(Paths.get(path)),
        e -> "Failed to read file: " + e.getMessage()
    );
}
```

### 2. Chaining Operations with andThen

```java
Result<User, String> result = validateUserId(userId)
    .andThen(id -> fetchUser(id))
    .andThen(user -> validatePermissions(user))
    .andThen(user -> updateUser(user));

// Early return on first error, propagates automatically
```

### 3. Transforming Success Values with map

```java
Result<Integer, String> age = parseAge("25");
Result<String, String> category = age.map(a -> {
    if (a < 18) return "minor";
    if (a < 65) return "adult";
    return "senior";
});
```

### 4. Transforming Error Values with mapErr

```java
Result<User, DatabaseError> dbResult = fetchUser(id);
Result<User, String> httpResult = dbResult.mapErr(err ->
    switch (err) {
        case NOT_FOUND -> "404: User not found";
        case CONNECTION_ERROR -> "500: Database connection failed";
        case TIMEOUT -> "504: Database timeout";
    }
);
```

### 5. Pattern Matching with match

```java
String message = parseAge(input).match(
    age -> "Valid age: " + age,
    error -> "Error: " + error
);
```

### 6. Side Effects with ifOk and ifErr

```java
parseAge(input)
    .ifOk(age -> logger.info("Parsed age: {}", age))
    .ifErr(error -> logger.error("Parse failed: {}", error));
```

### 7. Providing Defaults with unwrapOr

```java
int age = parseAge(input).unwrapOr(0);  // Default to 0 if parsing fails
```

### 8. Computing Defaults with unwrapOrElse

```java
int age = parseAge(input).unwrapOrElse(error -> {
    logger.warn("Parse error: {}, using default", error);
    return 18;  // Default adult age
});
```

## Naming Conventions

### Method Names

- **Returning Result**: Use verb names like `parse`, `validate`, `fetch`, `process`
  ```java
  Result<User, Error> fetchUser(String id);
  Result<Integer, ParseError> parseInt(String input);
  ```

- **Avoid "try" prefix**: The Result type already indicates fallibility
  ```java
  // Good
  Result<Config, Error> parseConfig(String json);

  // Avoid
  Result<Config, Error> tryParseConfig(String json);
  ```

### Error Type Names

- **Descriptive suffixes**: Use `Error`, `Failure`, or `Problem`
  ```java
  enum ValidationError { ... }
  sealed interface ParseFailure { ... }
  ```

- **Context-specific**: Include the domain or operation
  ```java
  enum DatabaseError { ... }
  enum NetworkError { ... }
  ```

## Migration Strategy

### Step 1: Identify Candidates

Look for methods that:
- Return a value and throw checked exceptions
- Have `throws` clauses with expected failures
- Use exceptions for control flow

### Step 2: Define Error Types

Before converting, define appropriate error enums or sealed interfaces.

### Step 3: Convert Method Signature

```java
// Before
public User getUser(String id) throws UserNotFoundException, DatabaseException {
    // ...
}

// After
public Result<User, UserError> getUser(String id) {
    // ...
}
```

### Step 4: Update Call Sites

```java
// Before
try {
    User user = getUser(id);
    processUser(user);
} catch (UserNotFoundException e) {
    return notFound();
} catch (DatabaseException e) {
    return serverError();
}

// After
return getUser(id).match(
    user -> { processUser(user); return ok(); },
    error -> switch (error) {
        case NOT_FOUND -> notFound();
        case DATABASE_ERROR -> serverError();
    }
);
```

### Step 5: Propagate Gradually

Start from leaf functions (lowest level) and propagate upward. This allows incremental adoption.

## Performance Considerations

### Memory Allocation

Result instances are allocated on each call. For hot paths:
- Consider caching common error instances
- Profile before optimizing
- The JVM's escape analysis often optimizes short-lived Results

### vs Exceptions

- **Result**: Constant-time allocation, no stack unwinding
- **Exceptions**: Slower due to stack trace capture, but only on error path
- **Rule of thumb**: If errors are common (>1% of calls), Result is faster

## Testing

### Testing Success Cases

```java
@Test
void testParseAge_validInput() {
    Result<Integer, ParseError> result = parseAge("25");
    assertTrue(result.isOk());
    assertEquals(25, result.unwrap());
}
```

### Testing Error Cases

```java
@Test
void testParseAge_invalidInput() {
    Result<Integer, ParseError> result = parseAge("abc");
    assertTrue(result.isErr());
    assertEquals(ParseError.INVALID_FORMAT, result.unwrapErr());
}
```

### Testing Chained Operations

```java
@Test
void testChaining_earlyError() {
    Result<String, String> result = parseAge("200")  // Out of range
        .andThen(age -> fetchUserByAge(age))
        .map(user -> user.getName());

    assertTrue(result.isErr());
    // fetchUserByAge should never be called
}
```

## Common Pitfalls

### 1. Unchecked unwrap()

**Bad**: Calling unwrap() without checking
```java
int age = parseAge(input).unwrap();  // May throw!
```

**Good**: Use safe methods or pattern matching
```java
int age = parseAge(input).unwrapOr(0);
// or
parseAge(input).match(
    value -> processAge(value),
    error -> handleError(error)
);
```

### 2. Mixing Exceptions and Results

**Bad**: Throwing exceptions from Result-returning methods
```java
Result<User, Error> getUser(String id) {
    if (id == null) {
        throw new IllegalArgumentException();  // Don't mix!
    }
    return fetchUser(id);
}
```

**Good**: Use Result consistently
```java
Result<User, Error> getUser(String id) {
    if (id == null) {
        return Result.err(Error.INVALID_ID);
    }
    return fetchUser(id);
}
```

### 3. Losing Error Information

**Bad**: Converting all errors to generic strings too early
```java
Result<User, String> result =
    fetchUser(id).mapErr(e -> "Error");  // Lost context!
```

**Good**: Preserve error information as long as possible
```java
Result<User, DatabaseError> result = fetchUser(id);
// Convert to string only at the boundary (e.g., HTTP response)
```

### 4. Not Using andThen for Chaining

**Bad**: Nested map calls creating Result<Result<T, E>, E>
```java
Result<Result<User, Error>, Error> nested =
    parseId(input).map(id -> fetchUser(id));  // Wrong!
```

**Good**: Use andThen for flattening
```java
Result<User, Error> result =
    parseId(input).andThen(id -> fetchUser(id));  // Correct!
```
