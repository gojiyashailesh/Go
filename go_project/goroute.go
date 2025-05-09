package main

import (
	"fmt"
	"time"
)

func printMessage(message string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(message, i)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	go printMessage("Goroutine")
	printMessage("Main function")
}
