package main

import (
	"fmt"
	"go_project/syntax"
)

func main() {
	fmt.Println("--------------------------------------------")
	fmt.Println("Above the output of the all the Syntax file ")
	fmt.Println("--------------------------------------------")
	syntax.Hello()
	fmt.Println("--------------------------------------------")
	syntax.ValueDemo()
	fmt.Println("--------------------------------------------")
	syntax.VariablesDemo()
	fmt.Println("------------------------------------------")
	syntax.LoopDemo()
	fmt.Println("------------------------------------------")
	syntax.IfElseDemo()
	fmt.Println("------------------------------------------")
	syntax.SwitchDemo()
	fmt.Println("------------------------------------------")
	syntax.ArraysDemo()
	fmt.Println("-----------------------------------------")
	syntax.SlicesDemo()
	fmt.Println("-----------------------------------------")
	syntax.MapsDemo()
	fmt.Println("----------------------------------------")
	syntax.Day1()
	fmt.Println("---------------------------------------")
	res := syntax.FunctionDemo()
	fmt.Println(res)
	fmt.Println("---------------------------------------")
	res2:= syntax.Func2Demo(44,55)
	fmt.Println(res2)
	fmt.Println("---------------------------------------") 
	x,y,z:=syntax.FuncWithMultiReturnValue(3,4,5)
	fmt.Println(x,y,z)
	fmt.Println("--------------------------------------")
	fmt.Println(syntax.VariableArgumentFunc(3,4,5,6,7,8,9))
	fmt.Println("--------------------------------------")
	fmt.Println(syntax.IntSeq())
	fmt.Println(syntax.IntSeq())
	fmt.Println("--------------------------------------")
	syntax.RangeFuncArr()
	fmt.Println("--------------------------------------")
	
}