package hellojava;

/**
 * Error-as-Value 패턴 데모
 *
 * 예외를 던지는 대신 Result 타입을 사용하여 에러를 값으로 다루는 방법을 보여줍니다.
 */
public class ErrorAsValueDemo {

    // 에러 타입 정의
    enum ParseError {
        INVALID_FORMAT,
        OUT_OF_RANGE,
        EMPTY_INPUT
    }

    enum ValidationError {
        INVALID_EMAIL,
        WEAK_PASSWORD,
        USERNAME_TOO_SHORT
    }

    /**
     * 예제 1: 숫자 파싱 (예외 대신 Result 사용)
     */
    public static Result<Integer, ParseError> parseAge(String input) {
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

    /**
     * 예제 2: 이메일 검증
     */
    public static Result<String, ValidationError> validateEmail(String email) {
        if (email == null || !email.contains("@")) {
            return Result.err(ValidationError.INVALID_EMAIL);
        }
        return Result.ok(email);
    }

    /**
     * 예제 3: 비밀번호 검증
     */
    public static Result<String, ValidationError> validatePassword(String password) {
        if (password == null || password.length() < 8) {
            return Result.err(ValidationError.WEAK_PASSWORD);
        }
        return Result.ok(password);
    }

    /**
     * 예제 4: 사용자명 검증
     */
    public static Result<String, ValidationError> validateUsername(String username) {
        if (username == null || username.length() < 3) {
            return Result.err(ValidationError.USERNAME_TOO_SHORT);
        }
        return Result.ok(username);
    }

    /**
     * 예제 5: 여러 검증을 체이닝
     */
    public static Result<UserData, String> validateAndCreateUser(
        String username,
        String email,
        String password
    ) {
        return validateUsername(username)
            .mapErr(e -> "Username error: " + e)
            .andThen(validUsername ->
                validateEmail(email)
                    .mapErr(e -> "Email error: " + e)
                    .andThen(validEmail ->
                        validatePassword(password)
                            .mapErr(e -> "Password error: " + e)
                            .map(validPassword ->
                                new UserData(validUsername, validEmail, validPassword)
                            )
                    )
            );
    }

    /**
     * 예제 6: Result를 사용한 안전한 나눗셈
     */
    public static Result<Double, String> safeDivide(double numerator, double denominator) {
        if (denominator == 0) {
            return Result.err("Division by zero");
        }
        return Result.ok(numerator / denominator);
    }

    /**
     * 사용자 데이터 레코드
     */
    record UserData(String username, String email, String password) {}

    public static void main(String[] args) {
        System.out.println("=== Error-as-Value 패턴 데모 ===\n");

        // 1. 숫자 파싱 예제
        System.out.println("1. 숫자 파싱:");
        demonstrateParseAge("25");
        demonstrateParseAge("200");
        demonstrateParseAge("abc");
        demonstrateParseAge("");

        // 2. 체이닝 예제
        System.out.println("\n2. 체이닝 (map과 andThen):");
        demonstrateChaining("25");
        demonstrateChaining("200");

        // 3. 사용자 검증 예제
        System.out.println("\n3. 사용자 생성 (여러 검증 체이닝):");
        demonstrateUserCreation("john", "john@example.com", "password123");
        demonstrateUserCreation("jo", "john@example.com", "password123");
        demonstrateUserCreation("john", "invalid-email", "password123");
        demonstrateUserCreation("john", "john@example.com", "pass");

        // 4. 안전한 나눗셈 예제
        System.out.println("\n4. 안전한 나눗셈:");
        demonstrateDivision(10.0, 2.0);
        demonstrateDivision(10.0, 0.0);

        // 5. Pattern Matching 예제
        System.out.println("\n5. Pattern Matching:");
        demonstratePatternMatching();

        System.out.println("\n=== 데모 완료 ===");
    }

    private static void demonstrateParseAge(String input) {
        Result<Integer, ParseError> result = parseAge(input);

        result.match(
            age -> {
                System.out.println("✅ 입력: '" + input + "' → 성공: " + age + "세");
                return null;
            },
            error -> {
                String errorMsg = switch (error) {
                    case INVALID_FORMAT -> "잘못된 형식";
                    case OUT_OF_RANGE -> "범위 초과 (0-150)";
                    case EMPTY_INPUT -> "빈 입력";
                };
                System.out.println("❌ 입력: '" + input + "' → 실패: " + errorMsg);
                return null;
            }
        );
    }

    private static void demonstrateChaining(String input) {
        Result<String, String> result = parseAge(input)
            .mapErr(e -> "파싱 에러: " + e)
            .map(age -> {
                if (age < 18) return "미성년자";
                if (age < 65) return "성인";
                return "노인";
            });

        result.ifOk(category -> System.out.println("✅ " + input + " → " + category))
              .ifErr(error -> System.out.println("❌ " + input + " → " + error));
    }

    private static void demonstrateUserCreation(
        String username,
        String email,
        String password
    ) {
        Result<UserData, String> result = validateAndCreateUser(username, email, password);

        result.match(
            user -> {
                System.out.println("✅ 사용자 생성 성공: " + user.username());
                return null;
            },
            error -> {
                System.out.println("❌ 사용자 생성 실패: " + error);
                return null;
            }
        );
    }

    private static void demonstrateDivision(double a, double b) {
        Result<Double, String> result = safeDivide(a, b);

        String message = result.match(
            value -> String.format("✅ %.1f ÷ %.1f = %.2f", a, b, value),
            error -> String.format("❌ %.1f ÷ %.1f → %s", a, b, error)
        );
        System.out.println(message);
    }

    private static void demonstratePatternMatching() {
        // unwrapOr: 실패 시 기본값 사용
        int age1 = parseAge("25").unwrapOr(0);
        System.out.println("parseAge('25').unwrapOr(0) = " + age1);

        int age2 = parseAge("abc").unwrapOr(0);
        System.out.println("parseAge('abc').unwrapOr(0) = " + age2);

        // unwrapOrElse: 실패 시 함수로 기본값 계산
        int age3 = parseAge("200").unwrapOrElse(error -> {
            System.out.println("  에러 발생: " + error + ", 기본값 18 사용");
            return 18;
        });
        System.out.println("parseAge('200').unwrapOrElse(...) = " + age3);

        // toOptional: Result를 Optional로 변환
        parseAge("30").toOptional().ifPresent(age ->
            System.out.println("Optional로 변환: " + age)
        );
    }
}
