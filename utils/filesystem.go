package utils

import (
	"bufio"
	"fmt"
	"os"
)

type FileService interface {
	FileExists(filePath string) bool
	DirExists(dirPath string) bool
	RemoveFile(dirPath string)
	CreateFile(dirPath string) *os.File
	WriteToFile(filePath string, content string)
	ReadLine(filePath string) string
	ReadFileToByteArray(filePath string) []byte
	WriteByteArrayToFile(b []byte, filePath string)
}

type fileService struct{}

func NewFileService() *fileService {
	return &fileService{}
}

func (f *fileService) FileExists(filePath string) bool {
	info, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func (f *fileService) DirExists(dirPath string) bool {
	info, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func (f *fileService) RemoveFile(dirPath string) {
	err := os.Remove(dirPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (f *fileService) CreateFile(dirPath string) *os.File {
	file, err := os.Create(dirPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return file
}

func (f *fileService) WriteToFile(filePath string, content string) {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		file = f.CreateFile(filePath)
	}

	defer file.Close()

	_, err = fmt.Fprint(file, content)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (f *fileService) ReadLine(filePath string) string {
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

func (f *fileService) ReadFileToByteArray(filePath string) []byte {
	b, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return b
}

func (f *fileService) WriteByteArrayToFile(b []byte, filePath string) {
	err := os.WriteFile(filePath, b, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
