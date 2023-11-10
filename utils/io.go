package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func FileExists(dirPath string) bool {
	info, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func CreateFile(dirPath string) {
	_, err := os.Create(dirPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Created file: ", dirPath)
}

func WriteToFile(filePath string, content string) {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	_, err = fmt.Fprintln(file, content)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Wrote to file: ", filePath)
}

func ReadLine(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var result string
	scanner.Scan()
	result = scanner.Text()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}
