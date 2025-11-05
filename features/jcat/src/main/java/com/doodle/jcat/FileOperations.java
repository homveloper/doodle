package com.doodle.jcat;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.nio.file.attribute.BasicFileAttributes;

/**
 * File operations using error-as-value pattern.
 */
public class FileOperations {

    /**
     * Reads entire file content as string.
     */
    public static Result<String, FileError> readFile(String filePath) {
        try {
            Path path = Paths.get(filePath);

            if (!Files.exists(path)) {
                return Result.err(new FileError.NotFound(filePath));
            }

            if (!Files.isRegularFile(path)) {
                return Result.err(new FileError.InvalidPath(filePath));
            }

            if (!Files.isReadable(path)) {
                return Result.err(new FileError.NotReadable(filePath));
            }

            String content = Files.readString(path);
            return Result.ok(content);

        } catch (IOException e) {
            return Result.err(new FileError.IoError(filePath, e.getMessage()));
        } catch (Exception e) {
            return Result.err(new FileError.InvalidPath(filePath));
        }
    }

    /**
     * Gets file metadata information.
     */
    public static Result<FileInfo, FileError> getFileInfo(String filePath) {
        try {
            Path path = Paths.get(filePath);

            if (!Files.exists(path)) {
                return Result.err(new FileError.NotFound(filePath));
            }

            BasicFileAttributes attrs = Files.readAttributes(path, BasicFileAttributes.class);

            FileInfo info = new FileInfo(
                path.toAbsolutePath().toString(),
                attrs.size(),
                attrs.isDirectory(),
                attrs.isRegularFile(),
                Files.isReadable(path),
                Files.isWritable(path),
                Files.isExecutable(path),
                attrs.lastModifiedTime()
            );

            return Result.ok(info);

        } catch (IOException e) {
            return Result.err(new FileError.IoError(filePath, e.getMessage()));
        } catch (Exception e) {
            return Result.err(new FileError.InvalidPath(filePath));
        }
    }
}
