package main

import "fmt"

func main() {
	var meyArray [3]int // # posições fixas
	meyArray[0] = 10
	meyArray[1] = 20
	meyArray[2] = 30

	fmt.Println(meyArray[len(meyArray)-1])

	for i, v := range meyArray {
		fmt.Printf("Index: %v, Value: %v\n", i, v)
	}
}
