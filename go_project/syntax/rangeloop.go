package syntax

import "fmt"

func RangeFuncArr() {
	num := []int{3, 4, 5, 6, 7, 8, 9}
	sum := 0

	// Summing all elements using range
	for _, n := range num {
		sum += n
	}
	fmt.Println("Sum:", sum)

	// This will print the index of the element if it matches 6
	for i, n := range num {
		if n == 6 {
			fmt.Println("Index of 6:", i)
		}
	}
}
