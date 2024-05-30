package main

import (
	"bufio"
	"fmt"
	"os"
)

func readLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func searchInText(lines []string, searchString string) []string {
	var results []string
	for _, line := range lines {
		if contains(line, searchString) {
			results = append(results, line)
		}
	}
	return results
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && s[0:len(substr)] == substr || len(s) > len(substr) && contains(s[1:], substr)
}

func main() {
	filePath := "text.txt" 
	lines, err := readLines(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("Text loaded from file:")
	for _, line := range lines {
		fmt.Println(line)
	}

	var searchString string
	fmt.Print("Enter search string: ")
	fmt.Scanln(&searchString)
	results := searchInText(lines, searchString)

	fmt.Println("Search results:")
	for _, result := range results {
		fmt.Println(result)
	}
}