package com.doodle.jcat;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class CommandParserTest {

    @Test
    void parse_shouldReturnHelpForEmptyArgs() {
        Command cmd = CommandParser.parse(new String[]{});
        assertInstanceOf(Command.Help.class, cmd);
    }

    @Test
    void parse_shouldReturnHelpForHelpCommand() {
        assertInstanceOf(Command.Help.class, CommandParser.parse(new String[]{"help"}));
        assertInstanceOf(Command.Help.class, CommandParser.parse(new String[]{"-h"}));
        assertInstanceOf(Command.Help.class, CommandParser.parse(new String[]{"--help"}));
    }

    @Test
    void parse_shouldReturnReadCommandWithFilePath() {
        Command cmd = CommandParser.parse(new String[]{"read", "test.txt"});
        assertInstanceOf(Command.Read.class, cmd);
        assertEquals("test.txt", ((Command.Read) cmd).filePath());
    }

    @Test
    void parse_shouldReturnInfoCommandWithFilePath() {
        Command cmd = CommandParser.parse(new String[]{"info", "test.txt"});
        assertInstanceOf(Command.Info.class, cmd);
        assertEquals("test.txt", ((Command.Info) cmd).filePath());
    }

    @Test
    void parse_shouldReturnInvalidForReadWithoutFilePath() {
        Command cmd = CommandParser.parse(new String[]{"read"});
        assertInstanceOf(Command.Invalid.class, cmd);
        assertTrue(((Command.Invalid) cmd).message().contains("file path"));
    }

    @Test
    void parse_shouldReturnInvalidForInfoWithoutFilePath() {
        Command cmd = CommandParser.parse(new String[]{"info"});
        assertInstanceOf(Command.Invalid.class, cmd);
        assertTrue(((Command.Invalid) cmd).message().contains("file path"));
    }

    @Test
    void parse_shouldReturnInvalidForUnknownCommand() {
        Command cmd = CommandParser.parse(new String[]{"unknown"});
        assertInstanceOf(Command.Invalid.class, cmd);
        assertTrue(((Command.Invalid) cmd).message().contains("Unknown"));
    }
}
