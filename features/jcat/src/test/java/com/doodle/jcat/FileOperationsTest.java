package com.doodle.jcat;

import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.io.TempDir;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;

import static org.junit.jupiter.api.Assertions.*;

class FileOperationsTest {

    @TempDir
    Path tempDir;

    @Test
    void readFile_shouldReturnContentForExistingFile() throws IOException {
        // Given
        Path testFile = tempDir.resolve("test.txt");
        String expectedContent = "Hello, World!\nThis is a test.";
        Files.writeString(testFile, expectedContent);

        // When
        Result<String, FileError> result = FileOperations.readFile(testFile.toString());

        // Then
        assertTrue(result.isOk());
        assertEquals(expectedContent, result.unwrap());
    }

    @Test
    void readFile_shouldReturnErrorForNonExistentFile() {
        // Given
        String nonExistentPath = tempDir.resolve("nonexistent.txt").toString();

        // When
        Result<String, FileError> result = FileOperations.readFile(nonExistentPath);

        // Then
        assertTrue(result.isErr());
        assertInstanceOf(FileError.NotFound.class, result.unwrapErr());
    }

    @Test
    void readFile_shouldReturnErrorForDirectory() throws IOException {
        // Given
        Path directory = tempDir.resolve("testdir");
        Files.createDirectory(directory);

        // When
        Result<String, FileError> result = FileOperations.readFile(directory.toString());

        // Then
        assertTrue(result.isErr());
        assertInstanceOf(FileError.InvalidPath.class, result.unwrapErr());
    }

    @Test
    void getFileInfo_shouldReturnInfoForExistingFile() throws IOException {
        // Given
        Path testFile = tempDir.resolve("test.txt");
        String content = "Test content";
        Files.writeString(testFile, content);

        // When
        Result<FileInfo, FileError> result = FileOperations.getFileInfo(testFile.toString());

        // Then
        assertTrue(result.isOk());
        FileInfo info = result.unwrap();
        assertEquals(content.length(), info.size());
        assertTrue(info.isRegularFile());
        assertFalse(info.isDirectory());
        assertTrue(info.isReadable());
    }

    @Test
    void getFileInfo_shouldReturnInfoForDirectory() throws IOException {
        // Given
        Path directory = tempDir.resolve("testdir");
        Files.createDirectory(directory);

        // When
        Result<FileInfo, FileError> result = FileOperations.getFileInfo(directory.toString());

        // Then
        assertTrue(result.isOk());
        FileInfo info = result.unwrap();
        assertTrue(info.isDirectory());
        assertFalse(info.isRegularFile());
    }

    @Test
    void getFileInfo_shouldReturnErrorForNonExistentFile() {
        // Given
        String nonExistentPath = tempDir.resolve("nonexistent.txt").toString();

        // When
        Result<FileInfo, FileError> result = FileOperations.getFileInfo(nonExistentPath);

        // Then
        assertTrue(result.isErr());
        assertInstanceOf(FileError.NotFound.class, result.unwrapErr());
    }

    @Test
    void fileInfo_shouldFormatSizeCorrectly() throws IOException {
        // Given
        Path smallFile = tempDir.resolve("small.txt");
        Path mediumFile = tempDir.resolve("medium.txt");

        Files.writeString(smallFile, "a".repeat(500));  // 500 bytes
        Files.writeString(mediumFile, "a".repeat(2048)); // 2 KB

        // When
        FileInfo smallInfo = FileOperations.getFileInfo(smallFile.toString()).unwrap();
        FileInfo mediumInfo = FileOperations.getFileInfo(mediumFile.toString()).unwrap();

        // Then
        assertTrue(smallInfo.formatSize().endsWith(" B"));
        assertTrue(mediumInfo.formatSize().endsWith(" KB"));
    }

    @Test
    void fileError_shouldProvideHumanReadableMessages() {
        // Given
        FileError notFound = new FileError.NotFound("/path/to/file.txt");
        FileError notReadable = new FileError.NotReadable("/path/to/file.txt");
        FileError ioError = new FileError.IoError("/path/to/file.txt", "Permission denied");
        FileError invalidPath = new FileError.InvalidPath("/path/to/file.txt");

        // Then
        assertTrue(notFound.message().contains("not found"));
        assertTrue(notReadable.message().contains("not readable"));
        assertTrue(ioError.message().contains("Permission denied"));
        assertTrue(invalidPath.message().contains("Invalid"));
    }

    @Test
    void result_shouldChainOperationsWithMap() throws IOException {
        // Given
        Path testFile = tempDir.resolve("test.txt");
        Files.writeString(testFile, "hello");

        // When
        Result<Integer, FileError> lengthResult = FileOperations.readFile(testFile.toString())
            .map(String::length);

        // Then
        assertTrue(lengthResult.isOk());
        assertEquals(5, lengthResult.unwrap());
    }

    @Test
    void result_shouldChainOperationsWithAndThen() throws IOException {
        // Given
        Path testFile = tempDir.resolve("test.txt");
        Files.writeString(testFile, testFile.toString());

        // When
        Result<FileInfo, FileError> result = FileOperations.readFile(testFile.toString())
            .andThen(content -> FileOperations.getFileInfo(content.trim()));

        // Then
        assertTrue(result.isOk());
    }

    @Test
    void result_shouldHandleMatchPattern() throws IOException {
        // Given
        Path testFile = tempDir.resolve("test.txt");
        Files.writeString(testFile, "content");

        // When
        String message = FileOperations.readFile(testFile.toString())
            .match(
                content -> "Success: " + content.length() + " chars",
                error -> "Error: " + error.message()
            );

        // Then
        assertEquals("Success: 7 chars", message);
    }
}
