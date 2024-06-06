package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type TextIndex struct {
	lines []string
	index map[string][]int
}

func NewTextIndex(lines []string) *TextIndex {
	ti := &TextIndex{
		lines: lines,
		index: make(map[string][]int),
	}
	ti.buildIndex()
	return ti
}

func (ti *TextIndex) buildIndex() {
	for i, line := range ti.lines {
		words := strings.Fields(line)
		for _, word := range words {
			word = strings.ToLower(strings.Trim(word, ".,!?\""))
			ti.index[word] = append(ti.index[word], i)
		}
	}
}

func (ti *TextIndex) Search(searchWord string) []string {
	searchWord = strings.ToLower(searchWord)
	lineNumbers, found := ti.index[searchWord]
	if !found {
		return []string{}
	}

	var results []string
	for _, lineNumber := range lineNumbers {
		results = append(results, ti.lines[lineNumber])
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

	textIndex := NewTextIndex(lines)

	fmt.Println("Text loaded from file:")
	for _, line := range textIndex.lines {
		fmt.Println(line)
	}

	var searchWord string
	fmt.Print("Enter search word: ")
	fmt.Scanln(&searchWord)
	results := textIndex.Search(searchWord)

	if len(results) > 0 {
		fmt.Println("Search results:")
		for _, result := range results {
			fmt.Println(result)
		}
	} else {
		fmt.Println("No matches found.")
	}
}