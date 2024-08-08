# Log Filter

## Description

This is a simple log filter that can be used to filter logs from a local system. It was created with Golang, and the objecite is to provide a simple way to filter logs via key words or regular expressions.

The filter is based on the following rules:

- The filter is not case insensitive.
- The software will try to match the key words in the log line. If the key word is found, the line will be saved in the output file.
- If the key word is not found, the line will be ignored.
- The application will return messages with the number of lines that were filtered, and the number of lines that were evaluated.

## Usage

To run with provided mockups:

```bash
go run . -inputFile mockups/to_the_moon.log -outputFileName to_the_moon_filtered.log -filterType string -filterDefinition launch
```

Using regex:

```bash
go run . -inputFile mockups/to_the_moon.log -outputFileName to_the_moon_filtered.log -filterType regex -filterDefinition "Apol+o"
```

## License

[MIT](LICENSE)
