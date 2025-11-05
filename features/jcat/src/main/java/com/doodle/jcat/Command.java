package com.doodle.jcat;

/**
 * Represents a CLI command.
 */
public sealed interface Command permits Command.Read, Command.Info, Command.Help, Command.Invalid {
    record Read(String filePath) implements Command {}
    record Info(String filePath) implements Command {}
    record Help() implements Command {}
    record Invalid(String message) implements Command {}
}
