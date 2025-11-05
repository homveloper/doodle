package com.doodle.jcat;

/**
 * JCat - A simple file information CLI tool.
 * All error handling uses the error-as-value pattern with Result<T, E>.
 */
public class JCat {

    private static final String VERSION = "1.0.0";
    private static final String HELP_TEXT = """
        JCat v%s - File Information CLI Tool

        Usage:
            jcat <command> <file>

        Commands:
            read <file>    Read and display file contents
            info <file>    Display file metadata information
            help           Show this help message

        Examples:
            jcat read ./sample.txt
            jcat info ./sample.txt

        All operations use error-as-value pattern for robust error handling.
        """;

    public static void main(String[] args) {
        Command cmd = CommandParser.parse(args);

        int exitCode = execute(cmd);
        System.exit(exitCode);
    }

    static int execute(Command cmd) {
        return switch (cmd) {
            case Command.Read(var filePath) -> executeRead(filePath);
            case Command.Info(var filePath) -> executeInfo(filePath);
            case Command.Help() -> executeHelp();
            case Command.Invalid(var message) -> {
                System.err.println("Error: " + message);
                System.err.println();
                executeHelp();
                yield 1;
            }
        };
    }

    private static int executeRead(String filePath) {
        Result<String, FileError> result = FileOperations.readFile(filePath);

        return result.match(
            content -> {
                System.out.print(content);
                return 0;
            },
            error -> {
                System.err.println("Error: " + error.message());
                return 1;
            }
        );
    }

    private static int executeInfo(String filePath) {
        Result<FileInfo, FileError> result = FileOperations.getFileInfo(filePath);

        return result.match(
            info -> {
                System.out.print(info.format());
                return 0;
            },
            error -> {
                System.err.println("Error: " + error.message());
                return 1;
            }
        );
    }

    private static int executeHelp() {
        System.out.println(String.format(HELP_TEXT, VERSION));
        return 0;
    }
}
