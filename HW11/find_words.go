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
)

func main() {

	wordRegex1 := regexp.MustCompile(`\b[АЕЄИІЇОУЮЯаеєиіїоуюя]\w*[бвгґджзйклмнпрстфхцчшщБВГҐДЖЗЙКЛМНПРСТФХЦЧШЩ]\b`)

	file, err := os.Open("text.txt")
	if err != nil {
		fmt.Println("Помилка відкриття файлу:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fmt.Println("Слова, що починаються на голосні та закінчуються на приголосні:")
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Читаємо рядок:", line) // Debugging output
		matches := wordRegex1.FindAllString(line, -1)
		if len(matches) == 0 {
			fmt.Println("Немає збігів у цьому рядку")
		}
		for _, match := range matches {
			fmt.Println("Знайдено слово:", match)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Помилка читання файлу:", err)
		return
	}

	file.Seek(0, 0) // Перемотати файл на початок
	scanner = bufio.NewScanner(file)
	fmt.Println("Слова, що складаються з двох однакових букв, розділених будь-яким символом:")
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Читаємо рядок:", line) // Debugging output
		// Custom logic to find words that consist of two identical letters separated by any character
		words := regexp.MustCompile(`\w+`).FindAllString(line, -1)
		for _, word := range words {
			for i := 0; i < len(word)-2; i++ {
				if word[i] == word[i+2] && word[i+1] != word[i] {
					fmt.Println("Знайдено слово:", word)
					break
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Помилка читання файлу:", err)
	}
}


