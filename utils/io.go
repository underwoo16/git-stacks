package utils

import (
	"bufio"
	"fmt"
	"os"
)

func FileExists(dirPath string) bool {
	info, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func RemoveFile(dirPath string) {
	err := os.Remove(dirPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func CreateFile(dirPath string) *os.File {
	file, err := os.Create(dirPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return file
}

func WriteToFile(filePath string, content string) {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		file = CreateFile(filePath)
	}

	defer file.Close()

	_, err = fmt.Fprint(file, content)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func ReadLine(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	result := scanner.Text()
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return result
}

func ReadFileToByteArray(filePath string) []byte {
	b, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return b
}

func WriteByteArrayToFile(b []byte, filePath string) {
	err := os.WriteFile(filePath, b, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
