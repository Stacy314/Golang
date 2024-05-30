package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func indexText(lines []string) map[string][]int {
	index := make(map[string][]int)
	for i, line := range lines {
		words := strings.Fields(line)
		for _, word := range words {
			word = strings.ToLower(strings.Trim(word, ".,!?\""))
			index[word] = append(index[word], i)
		}
	}
	return index
}


func searchByWord(index map[string][]int, lines []string, searchWord string) []string {
	searchWord = strings.ToLower(searchWord)
	lineNumbers, found := index[searchWord]
	if !found {
		return []string{}
	}

	var results []string
	for _, lineNumber := range lineNumbers {
		results = append(results, lines[lineNumber])
	}
	return results
}


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

func main() {
	filePath := "text.txt" 

	lines, err := readLines(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		fmt.Println("Ensure the file exists and the path is correct.")
		return
	}

	index := indexText(lines)

	fmt.Println("Text loaded from file:")
	for _, line := range lines {
		fmt.Println(line)
	}

	var searchWord string
	fmt.Print("Enter search word: ")
	fmt.Scanln(&searchWord)
	results := searchByWord(index, lines, searchWord)

	if len(results) > 0 {
		fmt.Println("Search results:")
		for _, result := range results {
			fmt.Println(result)
		}
	} else {
		fmt.Println("No matches found.")
	}
}