package syntax

import "fmt"

// ImediateFunc returns the result of an immediately invoked function.
func ImediateFunc() {
	result := func(a int, b int) int {
		return a * b
	}(4, 8)

	fmt.Println(result)
}
