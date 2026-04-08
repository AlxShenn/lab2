// go run task3.go
package main

import (
	"fmt"
	"math"
	"os"
)

// проверка на полный квадрат
func isPerfectSquare(x int) bool {
	root := int(math.Sqrt(float64(x)))
	return root*root == x
}

func main() {
	var choice int
	fmt.Println("1 - input from console\n2 - input from file")
	fmt.Scan(&choice)

	count := 0

	if choice == 1 {
		var n, x int
		fmt.Print("Enter N: ")
		fmt.Scan(&n)

		for i := 0; i < n; i++ {
			fmt.Scan(&x)
			if isPerfectSquare(x) {
				count++
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

		var n, x int
		fmt.Fscan(file, &n)

		for i := 0; i < n; i++ {
			fmt.Fscan(file, &x)
			if isPerfectSquare(x) {
				count++
			}
		}

	} else {
		fmt.Println("Invalid choice")
		return
	}

	fmt.Println("Result:", count)
}
