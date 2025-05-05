package syntax

import "fmt"

func AnonomousFunc() {
	add := func() string {
		return " response from the anonomous function"
	}
	fmt.Println(add())
}