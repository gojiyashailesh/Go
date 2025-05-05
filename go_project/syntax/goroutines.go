package syntax

import (
	"fmt"
	"time"
)

func greet(){
	fmt.Println("Wait is Over",time.Now())
}

func GorouteRunner(){
	go greet()
	fmt.Println("Wating",time.Now())
	time.Sleep(time.Minute)
}