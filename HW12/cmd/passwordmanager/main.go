/*Console password manager
Write a console program for storing passwords (a simplified analogue of 
the pass utility in UNIX). Password encryption is not implemented in this work.
Functionality:
* display names of saved passwords
* save password by name (password entry through fmt.Scan)
* get saved password
Additional conditions:
* we use tracer bullet development, that is, we write iteratively
* save the state in a file (so that passwords can be viewed between runs)
* use the recommended package structure (cmd, internal, ...)*/

package main

import (
	"flag"
	"fmt"
	"os"

	"HW13/internal/passwordmanager"
)

func main() {
	manager, err := passwordmanager.NewManager("passwords.json")
	if err != nil {
		fmt.Println("Помилка при завантаженні паролів:", err)
		return
	}

	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	putCmd := flag.NewFlagSet("put", flag.ExitOnError)
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)

	putName := putCmd.String("name", "", "Назва паролю")
	putPassword := putCmd.String("password", "", "Пароль")

	getName := getCmd.String("name", "", "Назва паролю")

	if len(os.Args) < 2 {
		fmt.Println("Не вказано команду. Використовуйте 'list', 'put', або 'get'.")
		return
	}

	switch os.Args[1] {
	case "list":
		listCmd.Parse(os.Args[2:])
		names := manager.ListNames()
		fmt.Println("Назви збережених паролів:")
		for _, name := range names {
			fmt.Println(name)
		}
	case "put":
		putCmd.Parse(os.Args[2:])
		if *putName == "" || *putPassword == "" {
			fmt.Println("Необхідно вказати назву та пароль.")
			return
		}
		if err := manager.SavePassword(*putName, *putPassword); err != nil {
			fmt.Println("Помилка при збереженні паролю:", err)
		} else {
			fmt.Println("Пароль успішно збережено")
		}
	case "get":
		getCmd.Parse(os.Args[2:])
		if *getName == "" {
			fmt.Println("Необхідно вказати назву паролю.")
			return
		}
		password, err := manager.GetPassword(*getName)
		if err != nil {
			fmt.Println("Помилка при отриманні паролю:", err)
		} else {
			fmt.Println("Пароль для", *getName, ":", password)
		}
	default:
		fmt.Println("Невідома команда. Використовуйте 'list', 'put', або 'get'.")
	}
}
