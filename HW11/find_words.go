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
    var content strings.Builder
    for scanner.Scan() {
        content.WriteString(scanner.Text() + " ")
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

    wordPattern := `\b\w+\b`
    re, err := regexp.Compile(wordPattern)
    if err != nil {
        fmt.Println("Error compiling regular expression:", err)
        return
    }

    words := re.FindAllString(content.String(), -1)

    checkPattern := func(word string) bool {
        length := len(word)
        for i := 0; i < length; i++ {
            for j := i + 1; j < length; j++ {
                if word[i] == word[j] {
                    return true
                }
            }
        }
        return false
    }

    matchingWords := []string{}
    for _, word := range words {
        if checkPattern(word) {
            matchingWords = append(matchingWords, word)
        }
    }

    if len(matchingWords) > 0 {
        fmt.Println("Found matching words:")
        for _, match := range matchingWords {
            fmt.Println(match)
        }
    } else {
        fmt.Println("No matching words found.")
    }
}