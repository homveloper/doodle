package com.doodle.jcat;

import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.io.TempDir;

import java.io.ByteArrayOutputStream;
import java.io.IOException;
import java.io.PrintStream;
import java.nio.file.Files;
import java.nio.file.Path;

import static org.junit.jupiter.api.Assertions.*;

class JCatTest {

    @TempDir
    Path tempDir;

    @Test
    void execute_shouldReturnZeroForHelpCommand() {
        int exitCode = JCat.execute(new Command.Help());
        assertEquals(0, exitCode);
    }

    @Test
    void execute_shouldReturnOneForInvalidCommand() {
        int exitCode = JCat.execute(new Command.Invalid("test error"));
        assertEquals(1, exitCode);
    }

    @Test
    void execute_shouldReadFileContentSuccessfully() throws IOException {
        // Given
        Path testFile = tempDir.resolve("test.txt");
        String content = "Hello, JCat!";
        Files.writeString(testFile, content);

        ByteArrayOutputStream outContent = new ByteArrayOutputStream();
        System.setOut(new PrintStream(outContent));

        // When
        int exitCode = JCat.execute(new Command.Read(testFile.toString()));

        // Then
        assertEquals(0, exitCode);
        assertEquals(content, outContent.toString());

        System.setOut(System.out);
    }

    @Test
    void execute_shouldReturnErrorForNonExistentFile() {
        ByteArrayOutputStream errContent = new ByteArrayOutputStream();
        System.setErr(new PrintStream(errContent));

        // When
        int exitCode = JCat.execute(new Command.Read("nonexistent.txt"));

        // Then
        assertEquals(1, exitCode);
        assertTrue(errContent.toString().contains("not found"));

        System.setErr(System.err);
    }

    @Test
    void execute_shouldShowFileInfoSuccessfully() throws IOException {
        // Given
        Path testFile = tempDir.resolve("test.txt");
        Files.writeString(testFile, "test content");

        ByteArrayOutputStream outContent = new ByteArrayOutputStream();
        System.setOut(new PrintStream(outContent));

        // When
        int exitCode = JCat.execute(new Command.Info(testFile.toString()));

        // Then
        assertEquals(0, exitCode);
        String output = outContent.toString();
        assertTrue(output.contains("Path:"));
        assertTrue(output.contains("Size:"));
        assertTrue(output.contains("Type:"));
        assertTrue(output.contains("Permissions:"));

        System.setOut(System.out);
    }
}
