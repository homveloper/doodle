package com.doodle.jcat;

/**
 * Parses command line arguments into Command objects.
 */
public class CommandParser {

    public static Command parse(String[] args) {
        if (args.length == 0) {
            return new Command.Help();
        }

        String subcommand = args[0];

        return switch (subcommand) {
            case "read" -> {
                if (args.length < 2) {
                    yield new Command.Invalid("'read' command requires a file path");
                }
                yield new Command.Read(args[1]);
            }
            case "info" -> {
                if (args.length < 2) {
                    yield new Command.Invalid("'info' command requires a file path");
                }
                yield new Command.Info(args[1]);
            }
            case "help", "-h", "--help" -> new Command.Help();
            default -> new Command.Invalid("Unknown command: " + subcommand);
        };
    }
}
