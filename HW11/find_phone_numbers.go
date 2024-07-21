/*1. Search for phone numbers in the contact data file. Task: Create a regular expression 
that can be used to find phone numbers written in different formats. For example, you might 
start with an expression that finds phone numbers that are 10 digits long, then expand it to 
add support for different formats, such as numbers with parentheses, spaces, and hyphens.*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	phoneRegex := regexp.MustCompile(`\(?\d{3}\)?[-.\s]?\d{3}[-.\s]?\d{4}`)

	file, err := os.Open("numbers.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fmt.Println("Found phone numbers:")
	for scanner.Scan() {
		line := scanner.Text()
		matches := phoneRegex.FindAllString(line, -1)
		for _, match := range matches {
			fmt.Println(match)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}