# GoSFV

GoSFV is a fast, simple, and OS-independent command-line tool designed for the swift verification of file integrity through *SimpleFileVerification*. Written entirely in Go, it offers reliability and speed in one package.

## Usage
After [installing Go](https://go.dev/dl/), you can easily compile the Go source code using the following command:
```bash
go build GoSFV.go
```

After compiling the file, you can use the resulting executable as follows:
```bash
GoSFV.exe PATH
```
Make sure to replace PATH with the actual path that contains a .sfv file and the files that need to be verified. GoSFV will report any verification failures and save the results in a text file named 'failedSFVs.txt,' located in the same directory as the .sfv file.

## Dependencies
GoSFV relies on the following Go packages, which are all part of the [standard library](https://pkg.go.dev/std):

- [bufio](https://pkg.go.dev/bufio@go1.21.0): Used for buffered I/O operations.
- [fmt](https://pkg.go.dev/fmt): Used for formatting input and output.
- [hash/crc32](https://pkg.go.dev/hash/crc32): Used for calculating CRC32 checksums.
- [os](https://pkg.go.dev/os): Used for working with the file system.
- [path/filepath](https://pkg.go.dev/path/filepath): Used for file path manipulation.
- [strings](https://pkg.go.dev/strings): Used for string manipulation.
