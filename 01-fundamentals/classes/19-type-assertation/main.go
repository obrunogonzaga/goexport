package main

import "fmt"

func main() {
	var minhaVar interface{} = "Bruno Gonzaga"
	println(minhaVar.(string))
	res, ok := minhaVar.(int)
	fmt.Printf("Value of res is %v and result of ok is: %v\n", res, ok)
}
