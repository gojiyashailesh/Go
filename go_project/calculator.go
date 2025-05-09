package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func additionNumber(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println("Result:", total)
	return total
}

func subtractionNumber(num1, num2 int) int {
	subresult := num1 - num2
	fmt.Println("Result:", subresult)
	return subresult
}

func multiplication(num1, num2 int) int {
	res := num1 * num2
	fmt.Println("Result:", res)
	return res
}

func divideNumber(num1, num2 int) (float64, error) {
	if num2 == 0 {
		return 0, errors.New("Divide by zero")
	}
	return float64(num1) / float64(num2), nil
}

func factorial(num int) int {
	if num == 0 || num == 1 {
		return 1
	}
	result := 1
	for i := 2; i <= num; i++ {
		result *= i
	}
	return result
}

func 

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n===== MENU =====")
		fmt.Println("1. Addition")
		fmt.Println("2. Subtraction")
		fmt.Println("3. Multiplication")
		fmt.Println("4. Division")
		fmt.Println("5. Factorial")
		fmt.Println("6. Exit")
		fmt.Print("Choose an option: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("Enter numbers separated by space: ")
			scanner.Scan()
			input := scanner.Text()
			numStrs := splitInput(input)
			nums := []int{}
			for _, str := range numStrs {
				n, err := strconv.Atoi(str)
				if err != nil {
					fmt.Println("Invalid input")
					continue
				}
				nums = append(nums, n)
			}
			additionNumber(nums...)

		case "2":
			num1, num2 := readTwoNumbers(scanner)
			subtractionNumber(num1, num2)

		case "3":
			num1, num2 := readTwoNumbers(scanner)
			multiplication(num1, num2)

		case "4":
			num1, num2 := readTwoNumbers(scanner)
			result, err := divideNumber(num1, num2)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Result:", result)
			}

		case "5":
			fmt.Print("Enter a number: ")
			scanner.Scan()
			numStr := scanner.Text()
			num, err := strconv.Atoi(numStr)
			if err != nil || num < 0 {
				fmt.Println("Please enter a valid non-negative integer")
				continue
			}
			fmt.Printf("Factorial of %d is: %d\n", num, factorial(num))

		case "6":
			fmt.Println("👋 Goodbye!")
			return

		default:
			fmt.Println("Invalid option. Try again.")
		}
	}
}

func splitInput(input string) []string {
	var parts []string
	curr := ""
	for _, ch := range input {
		if ch == ' ' {
			if curr != "" {
				parts = append(parts, curr)
				curr = ""
			}
		} else {
			curr += string(ch)
		}
	}
	if curr != "" {
		parts = append(parts, curr)
	}
	return parts
}

func readTwoNumbers(scanner *bufio.Scanner) (int, int) {
	fmt.Print("Enter first number: ")
	scanner.Scan()
	num1Str := scanner.Text()
	num1, err1 := strconv.Atoi(num1Str)

	fmt.Print("Enter second number: ")
	scanner.Scan()
	num2Str := scanner.Text()
	num2, err2 := strconv.Atoi(num2Str)

	if err1 != nil || err2 != nil {
		fmt.Println("Invalid input, treating as zero")
		return 0, 0
	}
	return num1, num2
}
