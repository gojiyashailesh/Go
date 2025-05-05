package syntax


//with this fuction as the argument we can modified the behaviours of the function
// this will provide the filter like facilities 
//easy to test


func FunctionAsArguement( a int , b int ,f func(int,int)int )int {
	return f(a,b)
}
