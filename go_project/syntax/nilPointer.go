package syntax

import "fmt"

func PointerExamples() {
	var ptr = new(int)
	*ptr = 10
	fmt.Println(*ptr)
}
