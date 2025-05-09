package syntax

import "fmt"

func LoopDemo() {
	//simple for loop
	for i := 10; i >= 0; i-- {
		fmt.Println(i)
	}

	//while looo style for loop
	i := 10
	j := 1
	for i >= 0 {
		fmt.Println(i * j)
		j++
		i--
	}

	//infinite for loop
	count := 5
	for {
		fmt.Println("sams")
		count--
		if count == 0 {
			break
		}
	}
	//for loop with range

	for n := range 10 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

}
