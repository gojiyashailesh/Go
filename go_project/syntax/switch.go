package syntax

import (
	"fmt"
	"time"
)

func SwitchDemo(){
	i:=3
	switch i{
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("Sorry")
	}

	//time lib in the go 

	switch time.Now().Weekday(){
	case time.Saturday,time.Sunday:
			fmt.Println("weekend enjoy")
	default:
			fmt.Println("Weekday work hard")
	}
}