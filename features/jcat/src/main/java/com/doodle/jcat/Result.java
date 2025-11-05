package com.doodle.jcat;

import java.util.function.Function;
import java.util.function.Consumer;

/**
 * Result type for error-as-value pattern.
 * Represents either success (Ok) or failure (Err).
 */
public sealed interface Result<T, E> permits Result.Ok, Result.Err {

    record Ok<T, E>(T value) implements Result<T, E> {}
    record Err<T, E>(E error) implements Result<T, E> {}

    static <T, E> Result<T, E> ok(T value) {
        return new Ok<>(value);
    }

    static <T, E> Result<T, E> err(E error) {
        return new Err<>(error);
    }

    default boolean isOk() {
        return this instanceof Ok;
    }

    default boolean isErr() {
        return this instanceof Err;
    }

    default T unwrap() {
        return switch (this) {
            case Ok(var value) -> value;
            case Err(var error) -> throw new RuntimeException("Called unwrap on Err: " + error);
        };
    }

    default T unwrapOr(T defaultValue) {
        return switch (this) {
            case Ok(var value) -> value;
            case Err(var error) -> defaultValue;
        };
    }

    default E unwrapErr() {
        return switch (this) {
            case Ok(var value) -> throw new RuntimeException("Called unwrapErr on Ok: " + value);
            case Err(var error) -> error;
        };
    }

    default <U> Result<U, E> map(Function<T, U> mapper) {
        return switch (this) {
            case Ok(var value) -> ok(mapper.apply(value));
            case Err(var error) -> err(error);
        };
    }

    default <U> Result<U, E> andThen(Function<T, Result<U, E>> mapper) {
        return switch (this) {
            case Ok(var value) -> mapper.apply(value);
            case Err(var error) -> err(error);
        };
    }

    default <U> U match(Function<T, U> okMapper, Function<E, U> errMapper) {
        return switch (this) {
            case Ok(var value) -> okMapper.apply(value);
            case Err(var error) -> errMapper.apply(error);
        };
    }

    default void ifOk(Consumer<T> consumer) {
        if (this instanceof Ok(var value)) {
            consumer.accept(value);
        }
    }

    default void ifErr(Consumer<E> consumer) {
        if (this instanceof Err(var error)) {
            consumer.accept(error);
        }
    }
}
