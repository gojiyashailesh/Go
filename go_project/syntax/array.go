package syntax

import "fmt"

func ArraysDemo() {
	var arr = [3]int{22, 33, 44}
	fmt.Println(arr)
	var arr2 [5]int
	fmt.Println(arr2)
	var arr3 = [6]int{2, 3, 4, 5}
	fmt.Println(arr3)

	//infer typed array in the golang
	arr4 := [...]int{2, 3, 3, 4, 4, 5, 6, 7, 7, 89, 66}
	fmt.Println(arr4)

	// Traversing arr4
	for i := 0; i < len(arr4); i++ {
		fmt.Println(arr4[i])
	}

	li := 0
	fi := len(arr4) - 1
	for li < fi {
		arr4[li], arr4[fi] = arr4[fi], arr4[li]
		li++
		fi--
	}
	fmt.Println("resvered: ", arr4)

	//2D array in the go
	var arr2d = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{6, 7, 8},
	}
	fmt.Println(arr2d)
	fmt.Println("Travese the #D arrays")
	for i := 0; i < len(arr2d); i++ {
		for j := 0; j < len(arr2d[i]); j++ {
			fmt.Printf("arr2d[%d][%d]=%d\n", i, j, arr2d[i][j])
		}
	}
}
