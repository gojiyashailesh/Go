package syntax

import "fmt"

func MapsDemo() {
	m := make(map[string]int)
	m["one"] = 12
	m["two"] = 14
	fmt.Println(m)
}
