package syntax

import "fmt"

func SlicesDemo() {
	var sls = []int{2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(sls)
	sls1 :=sls[1:4]
	fmt.Println(sls1)
	sls = append(sls, 88)
	fmt.Println(sls)
	

}