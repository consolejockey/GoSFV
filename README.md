# GoSFV

GoSFV is a fast and simple command-line tool designed for the swift verification of file integrity through Simple File Verification (*SFV*). Written entirely in Go, it offers reliability and speed in one package.

## Usage

You can use GoSFV in two ways:

1. **Drag and Drop**: Simply drag and drop the `.sfv` file or the folder containing the `.sfv` file onto the executable file (`gosfv.exe` on Windows, `gosfv` on Unix-like systems).

2. **Command Line**: You can start GoSFV from the command line by providing the path to the `.sfv` file as an argument: `gosfv <path to sfv file or folder>`

## Features

- Automatically detects the `.sfv` file in the provided directory.
- Verifies the checksums of files listed in the .sfv file.
- Saves a list of files with incorrect checksums to a separate text file (`failedSFVs.txt`) in the same directory as the .sfv file.

## Dependencies
GoSFV relies on the following Go packages, which are all part of the [standard library](https://pkg.go.dev/std):

- [bufio](https://pkg.go.dev/bufio@go1.21.0): Used for buffered I/O operations.
- [fmt](https://pkg.go.dev/fmt): Used for formatting input and output.
- [hash/crc32](https://pkg.go.dev/hash/crc32): Used for calculating CRC32 checksums.
- [os](https://pkg.go.dev/os): Used for working with the file system.
- [path/filepath](https://pkg.go.dev/path/filepath): Used for file path manipulation.
- [strings](https://pkg.go.dev/strings): Used for string manipulation.
