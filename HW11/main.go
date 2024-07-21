/*1. Search for phone numbers in the contact data file. Task: Create a regular expression 
that can be used to find phone numbers written in different formats. For example, you might 
start with an expression that finds phone numbers that are 10 digits long, then expand it to 
add support for different formats, such as numbers with parentheses, spaces, and hyphens.

2. Implement a search for words with a certain pattern in a text file. Task: Create a regular 
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
	// Задача 1: Пошук телефонних номерів
	phoneRegex := regexp.MustCompile(`\(?\d{3}\)?[-.\s]?\d{3}[-.\s]?\d{4}`)

	file, err := os.Open("numbers.txt")
	if err != nil {
		fmt.Println("Помилка відкриття файлу:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fmt.Println("Знайдені телефонні номери:")
	for scanner.Scan() {
		line := scanner.Text()
		matches := phoneRegex.FindAllString(line, -1)
		for _, match := range matches {
			fmt.Println(match)
		}
	}

	// Задача 2: Пошук слів з певним шаблоном
	wordRegex1 := regexp.MustCompile(`\b[АЕЄИІЇОУЮЯаеєиіїоуюя]\w*[бвгґджзйклмнпрстфхцчшщБВГҐДЖЗЙКЛМНПРСТФХЦЧШЩ]\b`)
	wordRegex2 := regexp.MustCompile(`\b(\w)\W?\1\b`)

	file, err = os.Open("text.txt")
	if err != nil {
		fmt.Println("Помилка відкриття файлу:", err)
		return
	}
	defer file.Close()

	scanner = bufio.NewScanner(file)
	fmt.Println("\nСлова, що починаються на голосні та закінчуються на приголосні:")
	for scanner.Scan() {
		line := scanner.Text()
		matches := wordRegex1.FindAllString(line, -1)
		for _, match := range matches {
			fmt.Println(match)
		}
	}

	file.Seek(0, 0) // Перемотати файл на початок
	scanner = bufio.NewScanner(file)
	fmt.Println("\nСлова, що складаються з двох однакових букв, розділених будь-яким символом:")
	for scanner.Scan() {
		line := scanner.Text()
		matches := wordRegex2.FindAllString(line, -1)
		for _, match := range matches {
			fmt.Println(match)
		}
	}
}