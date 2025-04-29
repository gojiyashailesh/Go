package syntax

import "fmt"

func VariablesDemo(){
	var name string = "Sams workspace"
	fmt.Println(name)
	var intdemo int = 12
	fmt.Println(intdemo)
	var flotdemo = 44.33
	fmt.Println(flotdemo)//by default it takes the float64


	//const 
	const username = "sams@ai"//not changeable
	fmt.Println(username)

	//shorthend 
	naming:=45000
	fmt.Println(naming)


	//multiple vars 
	var a,b int = 22,33
	fmt.Println(a,b)

	//multiple vars dtypes
	var (
		age     int  = 88
		showing bool = true
	)

	fmt.Println(age)
	fmt.Println(showing)


}