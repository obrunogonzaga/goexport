package main

import (
	"fmt"
	"os"
	"strconv"
)

// Write a function in Go that takes an integer as a parameter and doubles its value.
// Use pointers to modify the original variable.

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		panic("Please provide a single integer argument")
	}
	arg := args[0]

	num, err := strconv.Atoi(arg)
	if err != nil {
		panic("Please provide a single integer argument")
	}

	double(&num)
	fmt.Printf("The double of %d is %d\n", num/2, num)
}

func double(number *int) {
	*number *= 2
}
