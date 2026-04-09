package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isValidLocalPart(local string) bool {
	// Length check
	if len(local) < 6 || len(local) > 30 {
		return false
	}
	// Cannot start or end with a dot
	if local[0] == '.' || local[len(local)-1] == '.' {
		return false
	}
	// No consecutive dots
	if strings.Contains(local, "..") {
		return false
	}
	// Allowed characters: a-z, 0-9, ., +
	for _, c := range local {
		if !((c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') || c == '.' || c == '+') {
			return false
		}
	}
	return true
}

func normalize(email string) string {
	atPos := strings.Index(email, "@")
	if atPos <= 0 || atPos == len(email)-1 {
		return ""
	}
	local := email[:atPos]
	domain := email[atPos:]

	if !isValidLocalPart(local) {
		return ""
	}

	var newLocal strings.Builder
	for _, c := range local {
		if c == '+' {
			break
		}
		if c != '.' {
			newLocal.WriteRune(c)
		}
	}
	return newLocal.String() + domain
}

func main() {
	var choice int
	fmt.Println("1 - input from console\n2 - input from file")
	fmt.Scan(&choice)

	emails := make(map[string]bool)

	if choice == 1 {
		var n int
		fmt.Print("Enter number of emails: ")
		fmt.Scan(&n)

		for i := 0; i < n; i++ {
			var email string
			fmt.Scan(&email)
			if norm := normalize(email); norm != "" {
				emails[norm] = true
			}
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
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			email := scanner.Text()
			if norm := normalize(email); norm != "" {
				emails[norm] = true
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading file")
			return
		}
	} else {
		fmt.Println("Invalid choice")
		return
	}

	fmt.Println("Number of unique emails:", len(emails))
}
