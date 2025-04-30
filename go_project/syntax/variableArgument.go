package syntax

import "fmt"

func VariableArgumentFunc(sum...int)int{
	total:=0
	for _, num:= range sum {
		total +=num
	}
	fmt.Println(total)
	return total
}