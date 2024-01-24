package main

import "fmt"

// Work like generics in other languages, but now go have generics
func main() {
	var x interface{} = 10
	var y interface{} = "hello"

	showType(x)
	showType(y)
}

func showType(t interface{}) {
	fmt.Printf("Type is %T and value is %v\n", t, t)
}
