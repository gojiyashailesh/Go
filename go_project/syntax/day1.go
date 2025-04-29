package syntax

import (
	"fmt"
)

func Day1() {
	var name string = "shailesh"
	names := "shailesh"
	var age int = 19
	var weight float64 = 72.88
	const rights bool = true

	// Use format specifiers
	fmt.Printf("Name: %s\nNames: %s\nAge: %d\nWeight: %.4f\nRights: %t\n", name, names, age, weight, rights)

	//arrays
	// fixed in the size
	var myarr = [10]int{11, 12, 13, 14, 15, 16, 22, 23, 22}
	fmt.Println(myarr)

	//infer sized array
	//not fixed takes the infer size like variable arguments
	var myarr2 = [...]int{22, 33, 44, 55, 66}
	fmt.Println(myarr2)

	var myslice = []int{22, 22, 22, 22, 22, 22}
	fmt.Println(myslice)
	cslice := []float64{33, 33, 33, 33, 33, 33}
	fmt.Println(cslice)

	myslice = append(myslice, 66)
	//append data into the existing slice
	//slice is like the list of the python
	//array of the javascript
	cslice = append(cslice, 45)
	fmt.Println(cslice)

	makeslice := make([]string, 3, 5)
	makeslice = append(makeslice, "shailesh", "gojiya", "is the student")
	fmt.Println(makeslice)
	fmt.Println(len(makeslice))

	//struct and the type in the go
	//initalized the struct for the person type which contains the field like names age and balane
	type Person struct {
		name    string
		age     int
		balance float64
		data    complex64
	}

	//just like other langauge classes write the object
	var obj Person

	//initalize the value or say pass the values to the it
	obj.name = "shailesh gojiya"
	obj.age = 22
	obj.balance = 66.8888
	obj.data = complex64(3 + 2i)

	//print the output of the it
	fmt.Println(obj)

	//maps in the go
	var mymap map[string]int = make(map[string]int)
	mymap["shailesh"] = 28
	mymap["raju"] = 22
	mymap["gojiya"] = 99
	fmt.Println((mymap))

	//pointer -> hold the memory address of the other variables
	var expoint int = 22
	var p *int = &expoint
	fmt.Println(expoint)
	fmt.Println(*p)

	//typecasting in the goo
	var intTofloat int = 3000
	var result float64 = float64(intTofloat)
	fmt.Printf("./%.4f\n", result)

	//loop in the go
	//for loop
	var count int
	for count = 0; count <= 10; count++ {
		fmt.Print(count, "->")
	}

	//for loop 2
	var table int = 5
	for i := 0; i <= 10; i++ {
		fmt.Printf("%d * %d = %d\n", table, i, table*i)
	}

	//while loop
	i := 20
	for i > 0 {
		fmt.Print(i, " - ")
		i--
	}

	//for loop with range
	var arr = [...]int{2, 3, 4, 5, 6, 7, 8, 0}
	fmt.Println(len(arr))
	for i, val := range arr {
		fmt.Println(i, val)
	}

	//infinite loop
	// if i don't declare the break condition then this loop continue to work
	// it goes to the infinite don't stop until we break it or the kill the terminals
	counts := 0
	for {
		if counts == 5 {
			break
		}
		fmt.Println(counts)
		counts++
	}

	//conditional statements

	// if else
	var day string = "wed"

	if day == "monday" || day == "Tuesday" {
		fmt.Println("meeting")
	} else if day == "monday" && day == "tuesday" {
		fmt.Println("easy")
	} else if day == "saturday" {
		fmt.Println("enjoy")
	} else {
		fmt.Println("no in system")
	}

	//switch case
	choice := 5
	switch choice {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("Not present in the data")
	}

	//functions
	res := addTwo(3, 5)
	fmt.Println(res)
	res2 := variableArgument(2, 3, 4, 5, 6, 7, 8)
	fmt.Println(res2)

	// erruslt,res := multireturn(20, 0)
	// fmt.Println(erruslt)

}

// in the go function is treat as the value
/*
	In the Go Function Types are
	1.Simple Function
	2.Function with return type
	3.Function with multiple parameters
	4.Function with multiple return
	5.verdiac function
	6.Anonuoums function
	7.Directy use anonomous function
	8.function as the first class citinzen or pass the function to the another fucntion


*/
// functions
func addTwo(a, b int) int {
	return a + b
}

//variadic fucntions or say variable argument function

func variableArgument(sum ...int) int {
	total := 0
	for _, n := range sum {
		total += n
	}
	return total
}

// // function with multiple return statements
// func multireturn(a, b int) (int, error) {
// 	if b == 0 {
// 		return 0, errors.New("divide by zero error")
// 	}
// 	return a / b, nil
// }