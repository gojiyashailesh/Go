package syntax

import "fmt"

func IfElseDemo() {
	var weekends string = "s"

	if weekends == "sunday" {
		fmt.Println("You can rest today")
	} else {
		fmt.Println("No rest work hard")
	}
}
