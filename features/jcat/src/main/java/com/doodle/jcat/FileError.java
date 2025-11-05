package com.doodle.jcat;

/**
 * Error types for file operations.
 */
public sealed interface FileError {
    record NotFound(String path) implements FileError {}
    record NotReadable(String path) implements FileError {}
    record IoError(String path, String message) implements FileError {}
    record InvalidPath(String path) implements FileError {}

    default String message() {
        return switch (this) {
            case NotFound(var path) -> "File not found: " + path;
            case NotReadable(var path) -> "File not readable: " + path;
            case IoError(var path, var msg) -> "IO error reading " + path + ": " + msg;
            case InvalidPath(var path) -> "Invalid file path: " + path;
        };
    }
}
