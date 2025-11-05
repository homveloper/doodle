package com.doodle.jcat;

import java.nio.file.attribute.FileTime;
import java.time.ZoneId;
import java.time.format.DateTimeFormatter;

/**
 * File metadata information.
 */
public record FileInfo(
    String path,
    long size,
    boolean isDirectory,
    boolean isRegularFile,
    boolean isReadable,
    boolean isWritable,
    boolean isExecutable,
    FileTime lastModified
) {
    private static final DateTimeFormatter DATE_FORMATTER =
        DateTimeFormatter.ofPattern("yyyy-MM-dd HH:mm:ss");

    public String formatSize() {
        if (size < 1024) {
            return size + " B";
        } else if (size < 1024 * 1024) {
            return String.format("%.2f KB", size / 1024.0);
        } else if (size < 1024 * 1024 * 1024) {
            return String.format("%.2f MB", size / (1024.0 * 1024));
        } else {
            return String.format("%.2f GB", size / (1024.0 * 1024 * 1024));
        }
    }

    public String formatLastModified() {
        return lastModified.toInstant()
            .atZone(ZoneId.systemDefault())
            .format(DATE_FORMATTER);
    }

    public String format() {
        return String.format("""
            Path: %s
            Size: %s (%d bytes)
            Type: %s
            Permissions: %s%s%s
            Last Modified: %s
            """,
            path,
            formatSize(),
            size,
            isDirectory ? "Directory" : (isRegularFile ? "Regular File" : "Other"),
            isReadable ? "r" : "-",
            isWritable ? "w" : "-",
            isExecutable ? "x" : "-",
            formatLastModified()
        );
    }
}
