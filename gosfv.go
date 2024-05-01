package main

import (
	"bufio"
	"fmt"
	"hash/crc32"
	"os"
	"path/filepath"
	"strings"
)

const outputFileName = "failedSFVs.txt"

func decimalToHex(decimalChecksum uint32) string {
	hexChecksum := fmt.Sprintf("%08x", decimalChecksum)
	return hexChecksum
}

func calculateCRC32Checksum(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Failed to calculate checksum. Error: %v\n", err)
		return "", err
	}
	crcHash := crc32.ChecksumIEEE(data)
	hexCrcHash := decimalToHex(crcHash)
	return hexCrcHash, nil
}

func findSFVFileInDirectory(dirPath string) (string, error) {
	matches, err := filepath.Glob(filepath.Join(dirPath, "*.sfv"))
	if len(matches) == 0 {
		fmt.Printf("No .SFV file found at: %v\n", dirPath)
		return "", err
	}
	return matches[0], err
}

func verifySFV(dirPath string) ([]string, error) {
	var failedSFVs []string
	var sfvFilePath string
	var sfvFolder string
	var err error

	if strings.HasSuffix(dirPath, ".sfv") {
		sfvFilePath = dirPath
		sfvFolder = filepath.Dir(dirPath)
	} else {
		sfvFilePath, err = findSFVFileInDirectory(dirPath)
		if err != nil {
			return nil, err
		}
		sfvFolder = filepath.Dir(sfvFilePath)
	}

	file, err := os.Open(sfvFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		fileName := parts[0]
		expectedChecksum := strings.TrimSpace(parts[1])
		fullFilePath := filepath.Join(sfvFolder, fileName)

		calculatedChecksum, err := calculateCRC32Checksum(fullFilePath)
		if err != nil {
			failedSFV := fileName + " " + calculatedChecksum
			failedSFVs = append(failedSFVs, failedSFV)

		} else if strings.ToLower(expectedChecksum) == calculatedChecksum {
			continue

		} else {
			fmt.Printf("Failed SFV: %v\n", fileName)
			failedSFV := fileName + " " + calculatedChecksum
			failedSFVs = append(failedSFVs, failedSFV)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Failed calculating checksum. %v\n", err)
		return failedSFVs, err
	}

	if len(failedSFVs) == 0 {
		fmt.Println("All files successfully verified.")
		return failedSFVs, nil
	}

	outputFilePath := filepath.Join(sfvFolder, outputFileName)
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return failedSFVs, err
	}
	defer outputFile.Close()

	for _, verification := range failedSFVs {
		_, err := fmt.Fprintln(outputFile, verification)
		if err != nil {
			fmt.Println("Error writing to output file:", err)
			return failedSFVs, err
		}
	}

	fmt.Println("Saved list of corrupt files to:", outputFilePath)
	return failedSFVs, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: gosfv <path>")
		return
	}
	filePath := os.Args[1]
	verifySFV(filePath)
	fmt.Println("Press Enter to exit...")
	fmt.Scanln()
}
