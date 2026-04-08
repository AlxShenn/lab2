// go run task1.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

// Функция проверки и построения палиндрома
// Возвращает true, если палиндром можно составить, и сам палиндром
func canFormPalindrome(s string) (bool, string) {
	// Массив для подсчёта символов (ASCII)
	count := [256]int{}

	// Считаем, сколько раз встречается каждый символ
	for i := 0; i < len(s); i++ {
		count[s[i]]++
	}

	// Проверяем, сколько символов встречается нечётное количество раз
	oddCount := 0
	oddChar := byte(' ')

	for i := 0; i < 256; i++ {
		if count[i]%2 == 1 {
			oddCount++
			oddChar = byte(i)
		}
	}

	// Если нечётных символов больше одного - палиндром невозможен
	if oddCount > 1 {
		return false, ""
	}

	// Строим половину палиндрома
	half := make([]byte, 0)

	for i := 0; i < 256; i++ {
		if count[i] >= 2 {
			// Добавляем половину символов
			for j := 0; j < count[i]/2; j++ {
				half = append(half, byte(i))
			}
		}
	}

	// Сортируем для красивого вывода
	sort.Slice(half, func(a, b int) bool { return half[a] < half[b] })

	// Собираем палиндром: половина + центр + перевёрнутая половина
	palindrome := make([]byte, 0)
	palindrome = append(palindrome, half...)

	if oddCount == 1 {
		palindrome = append(palindrome, oddChar)
	}

	// Добавляем перевёрнутую половину
	for i := len(half) - 1; i >= 0; i-- {
		palindrome = append(palindrome, half[i])
	}

	return true, string(palindrome)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Select input method:")
	fmt.Println("1 - Input from console")
	fmt.Println("2 - Read from file")
	fmt.Print("Your choice: ")

	var choice int
	fmt.Scanln(&choice)

	if choice == 1 {
		// ========== ВВОД С КОНСОЛИ ==========
		fmt.Println("\n=== Console Input ===")
		fmt.Print("Enter a string: ")
		reader.ReadString('\n') // Очищаем буфер
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		possible, result := canFormPalindrome(input)
		if possible {
			fmt.Println("Result: Yes")
			fmt.Println("Palindrome:", result)
		} else {
			fmt.Println("Result: No")
		}
	} else if choice == 2 {
		// ========== ЧТЕНИЕ ИЗ ФАЙЛА ==========
		fmt.Println("\n=== File Input ===")
		fmt.Print("Enter filename (e.g., input.txt): ")
		var filename string
		fmt.Scanln(&filename)

		file, err := os.Open(filename)
		if err != nil {
			fmt.Println("Error: Could not open file", filename)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		lineNumber := 1

		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())

			// Пропускаем пустые строки
			if line == "" {
				continue
			}

			fmt.Printf("\nLine %d: \"%s\"\n", lineNumber, line)

			possible, result := canFormPalindrome(line)
			if possible {
				fmt.Println("  Result: Yes")
				fmt.Println("  Palindrome:", result)
			} else {
				fmt.Println("  Result: No")
			}

			lineNumber++
		}

		fmt.Printf("\nTotal lines processed: %d\n", lineNumber-1)
	} else {
		fmt.Println("Invalid choice!")
	}
}
