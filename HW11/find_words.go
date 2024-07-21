/*2. Implement a search for words with a certain pattern in a text file. Task: Create a regular 
expression that can be used to find words that match a certain pattern. For example, an expression 
that finds words that start with vowels and end with consonants, or words that consist of two 
identical letters separated by any character. The task is creative.*/

package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "strings"
)

func main() {
    file, err := os.Open("text.txt")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()
        words := strings.Fields(line)

        fmt.Println("Processing line:", line)

        fmt.Println("Words consisting of two identical letters separated by any character:")
        for _, word := range words {
            if hasIdenticalLetters(cleanWord(word)) {
                fmt.Println(cleanWord(word))
            }
        }
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading file:", err)
    }
}

func cleanWord(word string) string {
    re := regexp.MustCompile(`[^\p{L}]`)
    cleaned := re.ReplaceAllString(word, "")
    return strings.ToLower(cleaned)
}

func hasIdenticalLetters(word string) bool {
    for i := 0; i < len(word)-1; i++ {
        for j := i + 1; j < len(word); j++ {
            if word[i] == word[j] {
                return true
            }
        }
    }
    return false
}
