# Result Pattern Examples

This document contains practical examples of using the Result<T, E> pattern in real-world scenarios.

## Example 1: User Registration Service

```java
// Error types
enum RegistrationError {
    INVALID_EMAIL,
    WEAK_PASSWORD,
    USERNAME_TAKEN,
    DATABASE_ERROR
}

class UserService {
    private final UserRepository repository;
    private final EmailValidator emailValidator;

    public Result<User, RegistrationError> registerUser(
        String username,
        String email,
        String password
    ) {
        return validateEmail(email)
            .andThen(e -> validatePassword(password))
            .andThen(p -> checkUsernameAvailable(username))
            .andThen(available -> createUser(username, email, password))
            .andThen(user -> saveUser(user));
    }

    private Result<String, RegistrationError> validateEmail(String email) {
        if (!emailValidator.isValid(email)) {
            return Result.err(RegistrationError.INVALID_EMAIL);
        }
        return Result.ok(email);
    }

    private Result<String, RegistrationError> validatePassword(String password) {
        if (password.length() < 8) {
            return Result.err(RegistrationError.WEAK_PASSWORD);
        }
        return Result.ok(password);
    }

    private Result<Boolean, RegistrationError> checkUsernameAvailable(String username) {
        return Result.of(
            () -> !repository.existsByUsername(username),
            e -> RegistrationError.DATABASE_ERROR
        ).andThen(available -> {
            if (!available) {
                return Result.err(RegistrationError.USERNAME_TAKEN);
            }
            return Result.ok(true);
        });
    }

    private Result<User, RegistrationError> createUser(
        String username,
        String email,
        String password
    ) {
        User user = new User(username, email, hashPassword(password));
        return Result.ok(user);
    }

    private Result<User, RegistrationError> saveUser(User user) {
        return Result.of(
            () -> repository.save(user),
            e -> RegistrationError.DATABASE_ERROR
        );
    }

    private String hashPassword(String password) {
        // Password hashing logic
        return password; // Simplified
    }
}

// Usage in a controller
class UserController {
    private final UserService userService;

    public Response register(RegistrationRequest request) {
        return userService.registerUser(
            request.username(),
            request.email(),
            request.password()
        ).match(
            user -> Response.ok(new UserDto(user)),
            error -> switch (error) {
                case INVALID_EMAIL -> Response.badRequest("Invalid email address");
                case WEAK_PASSWORD -> Response.badRequest("Password must be at least 8 characters");
                case USERNAME_TAKEN -> Response.conflict("Username already taken");
                case DATABASE_ERROR -> Response.serverError("Database error occurred");
            }
        );
    }
}
```

## Example 2: Configuration Parser

```java
// Error types with rich information
sealed interface ConfigError {
    record FileNotFound(String path) implements ConfigError {}
    record InvalidJson(String content, String reason) implements ConfigError {}
    record MissingField(String fieldName) implements ConfigError {}
    record InvalidValue(String fieldName, String value, String expectedType) implements ConfigError {}
}

class ConfigParser {
    public Result<AppConfig, ConfigError> loadConfig(String path) {
        return readFile(path)
            .andThen(content -> parseJson(content))
            .andThen(json -> validateConfig(json))
            .andThen(json -> buildConfig(json));
    }

    private Result<String, ConfigError> readFile(String path) {
        return Result.of(
            () -> Files.readString(Paths.get(path)),
            e -> new ConfigError.FileNotFound(path)
        );
    }

    private Result<JsonObject, ConfigError> parseJson(String content) {
        return Result.of(
            () -> JsonParser.parse(content),
            e -> new ConfigError.InvalidJson(content, e.getMessage())
        );
    }

    private Result<JsonObject, ConfigError> validateConfig(JsonObject json) {
        if (!json.has("database")) {
            return Result.err(new ConfigError.MissingField("database"));
        }
        if (!json.has("port")) {
            return Result.err(new ConfigError.MissingField("port"));
        }
        return Result.ok(json);
    }

    private Result<AppConfig, ConfigError> buildConfig(JsonObject json) {
        return parsePort(json.get("port"))
            .andThen(port -> parseDatabaseConfig(json.get("database"))
                .map(db -> new AppConfig(port, db))
            );
    }

    private Result<Integer, ConfigError> parsePort(JsonElement element) {
        return Result.of(
            () -> element.getAsInt(),
            e -> new ConfigError.InvalidValue("port", element.toString(), "integer")
        ).andThen(port -> {
            if (port < 1 || port > 65535) {
                return Result.err(new ConfigError.InvalidValue(
                    "port",
                    String.valueOf(port),
                    "integer between 1 and 65535"
                ));
            }
            return Result.ok(port);
        });
    }

    private Result<DatabaseConfig, ConfigError> parseDatabaseConfig(JsonElement element) {
        // Similar validation logic
        return Result.ok(new DatabaseConfig(/* ... */));
    }
}

// Usage with detailed error reporting
class Application {
    public static void main(String[] args) {
        ConfigParser parser = new ConfigParser();
        parser.loadConfig("config.json")
            .ifOk(config -> startServer(config))
            .ifErr(error -> {
                String message = switch (error) {
                    case ConfigError.FileNotFound(var path) ->
                        "Configuration file not found: " + path;
                    case ConfigError.InvalidJson(var content, var reason) ->
                        "Invalid JSON: " + reason;
                    case ConfigError.MissingField(var field) ->
                        "Missing required field: " + field;
                    case ConfigError.InvalidValue(var field, var value, var expected) ->
                        "Invalid value for " + field + ": got '" + value +
                        "', expected " + expected;
                };
                System.err.println("Failed to load config: " + message);
                System.exit(1);
            });
    }

    private static void startServer(AppConfig config) {
        // Server startup logic
    }
}
```

## Example 3: HTTP Client with Retry Logic

```java
enum HttpError {
    NETWORK_ERROR,
    TIMEOUT,
    SERVER_ERROR,
    CLIENT_ERROR,
    PARSE_ERROR
}

class HttpClient {
    private final java.net.http.HttpClient client;
    private final int maxRetries;

    public HttpClient(int maxRetries) {
        this.client = java.net.http.HttpClient.newHttpClient();
        this.maxRetries = maxRetries;
    }

    public Result<String, HttpError> get(String url) {
        return getWithRetry(url, 0);
    }

    private Result<String, HttpError> getWithRetry(String url, int attempt) {
        Result<String, HttpError> result = performRequest(url);

        if (result.isErr() && attempt < maxRetries) {
            HttpError error = result.unwrapErr();
            if (shouldRetry(error)) {
                sleep(calculateBackoff(attempt));
                return getWithRetry(url, attempt + 1);
            }
        }

        return result;
    }

    private Result<String, HttpError> performRequest(String url) {
        HttpRequest request = HttpRequest.newBuilder()
            .uri(URI.create(url))
            .GET()
            .build();

        return Result.of(
            () -> {
                HttpResponse<String> response = client.send(
                    request,
                    HttpResponse.BodyHandlers.ofString()
                );
                return response;
            },
            e -> {
                if (e instanceof java.net.http.HttpTimeoutException) {
                    return HttpError.TIMEOUT;
                }
                return HttpError.NETWORK_ERROR;
            }
        ).andThen(response -> {
            int status = response.statusCode();
            if (status >= 500) {
                return Result.err(HttpError.SERVER_ERROR);
            }
            if (status >= 400) {
                return Result.err(HttpError.CLIENT_ERROR);
            }
            return Result.ok(response.body());
        });
    }

    private boolean shouldRetry(HttpError error) {
        return error == HttpError.TIMEOUT ||
               error == HttpError.NETWORK_ERROR ||
               error == HttpError.SERVER_ERROR;
    }

    private void sleep(long millis) {
        try {
            Thread.sleep(millis);
        } catch (InterruptedException e) {
            Thread.currentThread().interrupt();
        }
    }

    private long calculateBackoff(int attempt) {
        return (long) Math.pow(2, attempt) * 100; // Exponential backoff
    }
}

// Usage
class WeatherService {
    private final HttpClient httpClient;
    private final String apiKey;

    public Result<Weather, String> getCurrentWeather(String city) {
        String url = "https://api.weather.com/v1/current?city=" +
                     city + "&apiKey=" + apiKey;

        return httpClient.get(url)
            .mapErr(error -> switch (error) {
                case NETWORK_ERROR -> "Network connection failed";
                case TIMEOUT -> "Request timed out";
                case SERVER_ERROR -> "Weather service is down";
                case CLIENT_ERROR -> "Invalid city name";
                case PARSE_ERROR -> "Failed to parse response";
            })
            .andThen(json -> parseWeather(json));
    }

    private Result<Weather, String> parseWeather(String json) {
        return Result.of(
            () -> JsonParser.parseWeather(json),
            e -> "Failed to parse weather data: " + e.getMessage()
        );
    }
}
```

## Example 4: Database Transaction

```java
enum DbError {
    CONNECTION_FAILED,
    QUERY_FAILED,
    CONSTRAINT_VIOLATION,
    TRANSACTION_FAILED
}

class OrderRepository {
    private final DataSource dataSource;

    public Result<Order, DbError> createOrder(Order order, List<OrderItem> items) {
        return withTransaction(conn -> {
            return insertOrder(conn, order)
                .andThen(orderId -> insertOrderItems(conn, orderId, items))
                .andThen(itemIds -> updateInventory(conn, items))
                .map(success -> order.withId(orderId));
        });
    }

    private <T> Result<T, DbError> withTransaction(
        Function<Connection, Result<T, DbError>> operation
    ) {
        return getConnection().andThen(conn -> {
            try {
                conn.setAutoCommit(false);

                Result<T, DbError> result = operation.apply(conn);

                if (result.isOk()) {
                    conn.commit();
                } else {
                    conn.rollback();
                }

                conn.close();
                return result;

            } catch (SQLException e) {
                try { conn.rollback(); } catch (SQLException ignored) {}
                try { conn.close(); } catch (SQLException ignored) {}
                return Result.err(DbError.TRANSACTION_FAILED);
            }
        });
    }

    private Result<Connection, DbError> getConnection() {
        return Result.of(
            () -> dataSource.getConnection(),
            e -> DbError.CONNECTION_FAILED
        );
    }

    private Result<Long, DbError> insertOrder(Connection conn, Order order) {
        return Result.of(
            () -> {
                PreparedStatement stmt = conn.prepareStatement(
                    "INSERT INTO orders (customer_id, total) VALUES (?, ?)",
                    Statement.RETURN_GENERATED_KEYS
                );
                stmt.setLong(1, order.customerId());
                stmt.setBigDecimal(2, order.total());
                stmt.executeUpdate();

                ResultSet keys = stmt.getGeneratedKeys();
                if (keys.next()) {
                    return keys.getLong(1);
                }
                throw new SQLException("Failed to get order ID");
            },
            e -> {
                if (e.getMessage().contains("constraint")) {
                    return DbError.CONSTRAINT_VIOLATION;
                }
                return DbError.QUERY_FAILED;
            }
        );
    }

    private Result<List<Long>, DbError> insertOrderItems(
        Connection conn,
        long orderId,
        List<OrderItem> items
    ) {
        List<Long> ids = new ArrayList<>();
        for (OrderItem item : items) {
            Result<Long, DbError> result = insertOrderItem(conn, orderId, item);
            if (result.isErr()) {
                return Result.err(result.unwrapErr());
            }
            ids.add(result.unwrap());
        }
        return Result.ok(ids);
    }

    private Result<Long, DbError> insertOrderItem(
        Connection conn,
        long orderId,
        OrderItem item
    ) {
        return Result.of(
            () -> {
                PreparedStatement stmt = conn.prepareStatement(
                    "INSERT INTO order_items (order_id, product_id, quantity, price) " +
                    "VALUES (?, ?, ?, ?)",
                    Statement.RETURN_GENERATED_KEYS
                );
                stmt.setLong(1, orderId);
                stmt.setLong(2, item.productId());
                stmt.setInt(3, item.quantity());
                stmt.setBigDecimal(4, item.price());
                stmt.executeUpdate();

                ResultSet keys = stmt.getGeneratedKeys();
                if (keys.next()) {
                    return keys.getLong(1);
                }
                throw new SQLException("Failed to get order item ID");
            },
            e -> DbError.QUERY_FAILED
        );
    }

    private Result<Boolean, DbError> updateInventory(
        Connection conn,
        List<OrderItem> items
    ) {
        for (OrderItem item : items) {
            Result<Boolean, DbError> result = decrementStock(conn, item);
            if (result.isErr()) {
                return result;
            }
        }
        return Result.ok(true);
    }

    private Result<Boolean, DbError> decrementStock(Connection conn, OrderItem item) {
        return Result.of(
            () -> {
                PreparedStatement stmt = conn.prepareStatement(
                    "UPDATE inventory SET stock = stock - ? " +
                    "WHERE product_id = ? AND stock >= ?"
                );
                stmt.setInt(1, item.quantity());
                stmt.setLong(2, item.productId());
                stmt.setInt(3, item.quantity());

                int updated = stmt.executeUpdate();
                if (updated == 0) {
                    throw new SQLException("Insufficient stock");
                }
                return true;
            },
            e -> {
                if (e.getMessage().contains("Insufficient stock")) {
                    return DbError.CONSTRAINT_VIOLATION;
                }
                return DbError.QUERY_FAILED;
            }
        );
    }
}

// Usage in a service layer
class OrderService {
    private final OrderRepository repository;

    public Result<Order, String> placeOrder(Order order, List<OrderItem> items) {
        return repository.createOrder(order, items)
            .mapErr(error -> switch (error) {
                case CONNECTION_FAILED -> "Database connection failed. Please try again.";
                case QUERY_FAILED -> "Failed to create order. Please contact support.";
                case CONSTRAINT_VIOLATION -> "Some items are out of stock.";
                case TRANSACTION_FAILED -> "Transaction failed. Order was not created.";
            });
    }
}
```

## Example 5: Multi-Step Validation Pipeline

```java
sealed interface ValidationError {
    record Required(String field) implements ValidationError {}
    record TooShort(String field, int min, int actual) implements ValidationError {}
    record TooLong(String field, int max, int actual) implements ValidationError {}
    record InvalidFormat(String field, String pattern) implements ValidationError {}
    record Custom(String message) implements ValidationError {}
}

class Validator<T> {
    private final T value;
    private final List<ValidationError> errors = new ArrayList<>();

    private Validator(T value) {
        this.value = value;
    }

    public static <T> Validator<T> of(T value) {
        return new Validator<>(value);
    }

    public Validator<T> required(String field) {
        if (value == null || (value instanceof String s && s.isEmpty())) {
            errors.add(new ValidationError.Required(field));
        }
        return this;
    }

    public Validator<T> minLength(String field, int min) {
        if (value instanceof String s && s.length() < min) {
            errors.add(new ValidationError.TooShort(field, min, s.length()));
        }
        return this;
    }

    public Validator<T> maxLength(String field, int max) {
        if (value instanceof String s && s.length() > max) {
            errors.add(new ValidationError.TooLong(field, max, s.length()));
        }
        return this;
    }

    public Validator<T> matches(String field, String pattern) {
        if (value instanceof String s && !s.matches(pattern)) {
            errors.add(new ValidationError.InvalidFormat(field, pattern));
        }
        return this;
    }

    public Validator<T> custom(boolean condition, String message) {
        if (!condition) {
            errors.add(new ValidationError.Custom(message));
        }
        return this;
    }

    public Result<T, List<ValidationError>> validate() {
        if (errors.isEmpty()) {
            return Result.ok(value);
        }
        return Result.err(new ArrayList<>(errors));
    }
}

// Usage
record UserInput(String username, String email, String password) {
    public Result<UserInput, List<ValidationError>> validate() {
        var usernameResult = Validator.of(username)
            .required("username")
            .minLength("username", 3)
            .maxLength("username", 20)
            .matches("username", "^[a-zA-Z0-9_]+$")
            .validate();

        var emailResult = Validator.of(email)
            .required("email")
            .matches("email", "^[A-Za-z0-9+_.-]+@[A-Za-z0-9.-]+$")
            .validate();

        var passwordResult = Validator.of(password)
            .required("password")
            .minLength("password", 8)
            .custom(
                password.matches(".*[A-Z].*"),
                "Password must contain at least one uppercase letter"
            )
            .custom(
                password.matches(".*[0-9].*"),
                "Password must contain at least one digit"
            )
            .validate();

        // Combine all validation results
        if (usernameResult.isErr() || emailResult.isErr() || passwordResult.isErr()) {
            List<ValidationError> allErrors = new ArrayList<>();
            usernameResult.ifErr(allErrors::addAll);
            emailResult.ifErr(allErrors::addAll);
            passwordResult.ifErr(allErrors::addAll);
            return Result.err(allErrors);
        }

        return Result.ok(this);
    }
}

// Controller usage
class SignupController {
    public Response signup(UserInput input) {
        return input.validate().match(
            validInput -> {
                // Proceed with registration
                return Response.ok("Registration successful");
            },
            errors -> {
                // Format errors for response
                Map<String, List<String>> errorMap = new HashMap<>();
                for (ValidationError error : errors) {
                    String message = formatError(error);
                    String field = extractField(error);
                    errorMap.computeIfAbsent(field, k -> new ArrayList<>()).add(message);
                }
                return Response.badRequest(errorMap);
            }
        );
    }

    private String formatError(ValidationError error) {
        return switch (error) {
            case ValidationError.Required(var field) ->
                field + " is required";
            case ValidationError.TooShort(var field, var min, var actual) ->
                field + " must be at least " + min + " characters";
            case ValidationError.TooLong(var field, var max, var actual) ->
                field + " must be at most " + max + " characters";
            case ValidationError.InvalidFormat(var field, var pattern) ->
                field + " has invalid format";
            case ValidationError.Custom(var message) ->
                message;
        };
    }

    private String extractField(ValidationError error) {
        return switch (error) {
            case ValidationError.Required(var field) -> field;
            case ValidationError.TooShort(var field, var min, var actual) -> field;
            case ValidationError.TooLong(var field, var max, var actual) -> field;
            case ValidationError.InvalidFormat(var field, var pattern) -> field;
            case ValidationError.Custom(var message) -> "general";
        };
    }
}
```

## Example 6: File Processing Pipeline

```java
enum FileError {
    NOT_FOUND,
    NOT_READABLE,
    INVALID_FORMAT,
    PROCESSING_FAILED
}

class CsvProcessor {
    public Result<Report, FileError> processFile(Path filePath) {
        return validateFile(filePath)
            .andThen(path -> readLines(path))
            .andThen(lines -> parseRecords(lines))
            .andThen(records -> validateRecords(records))
            .andThen(records -> processRecords(records))
            .map(data -> generateReport(data));
    }

    private Result<Path, FileError> validateFile(Path path) {
        if (!Files.exists(path)) {
            return Result.err(FileError.NOT_FOUND);
        }
        if (!Files.isReadable(path)) {
            return Result.err(FileError.NOT_READABLE);
        }
        return Result.ok(path);
    }

    private Result<List<String>, FileError> readLines(Path path) {
        return Result.of(
            () -> Files.readAllLines(path),
            e -> FileError.NOT_READABLE
        );
    }

    private Result<List<Record>, FileError> parseRecords(List<String> lines) {
        List<Record> records = new ArrayList<>();
        for (int i = 1; i < lines.size(); i++) { // Skip header
            Result<Record, FileError> recordResult = parseLine(lines.get(i), i);
            if (recordResult.isErr()) {
                return Result.err(recordResult.unwrapErr());
            }
            records.add(recordResult.unwrap());
        }
        return Result.ok(records);
    }

    private Result<Record, FileError> parseLine(String line, int lineNumber) {
        return Result.of(
            () -> {
                String[] parts = line.split(",");
                if (parts.length != 3) {
                    throw new IllegalArgumentException(
                        "Line " + lineNumber + ": expected 3 fields, got " + parts.length
                    );
                }
                return new Record(parts[0], parts[1], parts[2]);
            },
            e -> FileError.INVALID_FORMAT
        );
    }

    private Result<List<Record>, FileError> validateRecords(List<Record> records) {
        for (Record record : records) {
            if (!isValidRecord(record)) {
                return Result.err(FileError.INVALID_FORMAT);
            }
        }
        return Result.ok(records);
    }

    private boolean isValidRecord(Record record) {
        return record.name() != null &&
               record.value() != null &&
               !record.name().isEmpty();
    }

    private Result<ProcessedData, FileError> processRecords(List<Record> records) {
        return Result.of(
            () -> {
                // Complex processing logic
                ProcessedData data = new ProcessedData();
                for (Record record : records) {
                    data.add(transform(record));
                }
                return data;
            },
            e -> FileError.PROCESSING_FAILED
        );
    }

    private TransformedRecord transform(Record record) {
        // Transformation logic
        return new TransformedRecord(record.name(), record.value());
    }

    private Report generateReport(ProcessedData data) {
        return new Report(data.summary(), data.details());
    }
}

// Usage with detailed error handling
class FileProcessingService {
    private final CsvProcessor processor;

    public void processFile(String filePath) {
        processor.processFile(Paths.get(filePath))
            .ifOk(report -> {
                System.out.println("Processing successful!");
                System.out.println(report);
                saveReport(report);
            })
            .ifErr(error -> {
                String message = switch (error) {
                    case NOT_FOUND -> "File not found: " + filePath;
                    case NOT_READABLE -> "Cannot read file: " + filePath;
                    case INVALID_FORMAT -> "Invalid CSV format in file";
                    case PROCESSING_FAILED -> "Failed to process records";
                };
                System.err.println("Error: " + message);
                logError(message, error);
            });
    }

    private void saveReport(Report report) {
        // Save logic
    }

    private void logError(String message, FileError error) {
        // Logging logic
    }
}
```
