package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"HW13/internal/passwordmanager"
)

func main() {
	manager, err := passwordmanager.NewManager("passwords.json")
	if err != nil {
		fmt.Println("Помилка при завантаженні паролів:", err)
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\nВиберіть дію:")
		fmt.Println("1. Вивести назви збережених паролів")
		fmt.Println("2. Зберегти новий пароль")
		fmt.Println("3. Дістати збережений пароль")
		fmt.Println("4. Вихід")
		fmt.Print("Ваш вибір: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Println("Назви збережених паролів:")
			for _, name := range manager.ListNames() {
				fmt.Println(name)
			}
		case "2":
			fmt.Print("Введіть назву паролю: ")
			scanner.Scan()
			name := scanner.Text()

			fmt.Print("Введіть пароль: ")
			scanner.Scan()
			password := scanner.Text()

			if err := manager.SavePassword(name, password); err != nil {
				fmt.Println("Помилка при збереженні паролю:", err)
			} else {
				fmt.Println("Пароль успішно збережено")
			}
		case "3":
			fmt.Print("Введіть назву паролю: ")
			scanner.Scan()
			name := scanner.Text()

			password, err := manager.GetPassword(name)
			if err != nil {
				fmt.Println("Помилка при отриманні паролю:", err)
			} else {
				fmt.Println("Пароль для", name, ":", password)
			}
		case "4":
			return
		default:
			fmt.Println("Невідомий вибір, спробуйте ще раз.")
		}
	}
}