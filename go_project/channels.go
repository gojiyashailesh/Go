package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		ch <- "Input in the channnels" //this is data passed into the channels
	}()
	msg := <-ch //main channelt taking the data from the channels
	fmt.Println(msg)
}
