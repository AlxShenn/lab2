// go run task2.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// функция для "нормализации" email
func normalize(email string) string {
	atPos := strings.Index(email, "@")
	name := email[:atPos]
	domain := email[atPos:]

	newName := ""
	for i := 0; i < len(name); i++ {
		if name[i] == '+' {
			break // игнорируем всё после +
		}
		if name[i] != '.' {
			newName += string(name[i])
		}
	}

	return newName + domain
}

func main() {
	emails := make(map[string]bool)
	var choice int

	fmt.Println("1 - input from console\n2 - input from file")
	fmt.Scan(&choice)

	if choice == 1 {
		var n int
		fmt.Print("Enter number of emails: ")
		fmt.Scan(&n)

		var email string
		for i := 0; i < n; i++ {
			fmt.Scan(&email)
			emails[normalize(email)] = true
		}
	} else if choice == 2 {
		var filename string
		fmt.Print("Enter file name: ")
		fmt.Scan(&filename)

		file, err := os.Open(filename)
		if err != nil {
			fmt.Println("File open error")
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			email := scanner.Text()
			emails[normalize(email)] = true
		}
	} else {
		fmt.Println("Invalid choice")
		return
	}

	fmt.Println("Number of unique emails:", len(emails))
}
