package hellojava;

import java.util.Objects;
import java.util.Optional;
import java.util.function.Consumer;
import java.util.function.Function;

/**
 * A Result type that represents either success (Ok) or failure (Err).
 * This enables error-as-value pattern in Java, avoiding exception throwing.
 *
 * @param <T> The type of the success value
 * @param <E> The type of the error value
 */
public sealed interface Result<T, E> permits Result.Ok, Result.Err {

    /**
     * Creates a successful Result containing a value.
     *
     * @param value The success value
     * @param <T> The type of the success value
     * @param <E> The type of the error value
     * @return A Result containing the success value
     */
    static <T, E> Result<T, E> ok(T value) {
        return new Ok<>(value);
    }

    /**
     * Creates a failed Result containing an error.
     *
     * @param error The error value
     * @param <T> The type of the success value
     * @param <E> The type of the error value
     * @return A Result containing the error value
     */
    static <T, E> Result<T, E> err(E error) {
        return new Err<>(error);
    }

    /**
     * Wraps a potentially exception-throwing operation in a Result.
     *
     * @param supplier The operation that might throw
     * @param errorMapper Function to convert Exception to error type
     * @param <T> The type of the success value
     * @param <E> The type of the error value
     * @return A Result containing either the success value or mapped error
     */
    static <T, E> Result<T, E> of(ThrowingSupplier<T> supplier, Function<Exception, E> errorMapper) {
        try {
            return ok(supplier.get());
        } catch (Exception e) {
            return err(errorMapper.apply(e));
        }
    }

    /**
     * Returns true if this Result is Ok.
     */
    boolean isOk();

    /**
     * Returns true if this Result is Err.
     */
    boolean isErr();

    /**
     * Returns the success value if Ok, or throws if Err.
     *
     * @throws IllegalStateException if this Result is Err
     */
    T unwrap();

    /**
     * Returns the success value if Ok, or the provided default if Err.
     */
    T unwrapOr(T defaultValue);

    /**
     * Returns the success value if Ok, or computes a default from the error if Err.
     */
    T unwrapOrElse(Function<E, T> fn);

    /**
     * Returns the error value if Err, or throws if Ok.
     *
     * @throws IllegalStateException if this Result is Ok
     */
    E unwrapErr();

    /**
     * Maps the success value using the provided function.
     * If this Result is Err, returns the error unchanged.
     */
    <U> Result<U, E> map(Function<T, U> fn);

    /**
     * Maps the error value using the provided function.
     * If this Result is Ok, returns the success value unchanged.
     */
    <F> Result<T, F> mapErr(Function<E, F> fn);

    /**
     * Applies a function that returns a Result to the success value.
     * This is used for chaining operations that may fail.
     * Also known as flatMap.
     */
    <U> Result<U, E> andThen(Function<T, Result<U, E>> fn);

    /**
     * Executes the provided consumer if this Result is Ok.
     */
    Result<T, E> ifOk(Consumer<T> consumer);

    /**
     * Executes the provided consumer if this Result is Err.
     */
    Result<T, E> ifErr(Consumer<E> consumer);

    /**
     * Converts this Result to an Optional.
     * Returns Optional.of(value) if Ok, Optional.empty() if Err.
     */
    Optional<T> toOptional();

    /**
     * Pattern matching for Result.
     * Executes onOk if this is Ok, onErr if this is Err.
     *
     * @param onOk Function to apply to success value
     * @param onErr Function to apply to error value
     * @return The result of applying the appropriate function
     */
    <U> U match(Function<T, U> onOk, Function<E, U> onErr);

    /**
     * Ok variant of Result containing a success value.
     */
    record Ok<T, E>(T value) implements Result<T, E> {
        public Ok {
            Objects.requireNonNull(value, "Ok value cannot be null");
        }

        @Override
        public boolean isOk() {
            return true;
        }

        @Override
        public boolean isErr() {
            return false;
        }

        @Override
        public T unwrap() {
            return value;
        }

        @Override
        public T unwrapOr(T defaultValue) {
            return value;
        }

        @Override
        public T unwrapOrElse(Function<E, T> fn) {
            return value;
        }

        @Override
        public E unwrapErr() {
            throw new IllegalStateException("Called unwrapErr on Ok value: " + value);
        }

        @Override
        public <U> Result<U, E> map(Function<T, U> fn) {
            return ok(fn.apply(value));
        }

        @Override
        public <F> Result<T, F> mapErr(Function<E, F> fn) {
            return ok(value);
        }

        @Override
        public <U> Result<U, E> andThen(Function<T, Result<U, E>> fn) {
            return fn.apply(value);
        }

        @Override
        public Result<T, E> ifOk(Consumer<T> consumer) {
            consumer.accept(value);
            return this;
        }

        @Override
        public Result<T, E> ifErr(Consumer<E> consumer) {
            return this;
        }

        @Override
        public Optional<T> toOptional() {
            return Optional.of(value);
        }

        @Override
        public <U> U match(Function<T, U> onOk, Function<E, U> onErr) {
            return onOk.apply(value);
        }
    }

    /**
     * Err variant of Result containing an error value.
     */
    record Err<T, E>(E error) implements Result<T, E> {
        public Err {
            Objects.requireNonNull(error, "Err value cannot be null");
        }

        @Override
        public boolean isOk() {
            return false;
        }

        @Override
        public boolean isErr() {
            return true;
        }

        @Override
        public T unwrap() {
            throw new IllegalStateException("Called unwrap on Err value: " + error);
        }

        @Override
        public T unwrapOr(T defaultValue) {
            return defaultValue;
        }

        @Override
        public T unwrapOrElse(Function<E, T> fn) {
            return fn.apply(error);
        }

        @Override
        public E unwrapErr() {
            return error;
        }

        @Override
        public <U> Result<U, E> map(Function<T, U> fn) {
            return err(error);
        }

        @Override
        public <F> Result<T, F> mapErr(Function<E, F> fn) {
            return err(fn.apply(error));
        }

        @Override
        public <U> Result<U, E> andThen(Function<T, Result<U, E>> fn) {
            return err(error);
        }

        @Override
        public Result<T, E> ifOk(Consumer<T> consumer) {
            return this;
        }

        @Override
        public Result<T, E> ifErr(Consumer<E> consumer) {
            consumer.accept(error);
            return this;
        }

        @Override
        public Optional<T> toOptional() {
            return Optional.empty();
        }

        @Override
        public <U> U match(Function<T, U> onOk, Function<E, U> onErr) {
            return onErr.apply(error);
        }
    }

    /**
     * Functional interface for operations that may throw exceptions.
     */
    @FunctionalInterface
    interface ThrowingSupplier<T> {
        T get() throws Exception;
    }
}
