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
	CheckError(err)
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
	CheckError(err)
}

func ReadLine(filePath string) string {
	file, err := os.Open(filePath)
	CheckError(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	result := scanner.Text()
	CheckError(scanner.Err())

	return result
}
