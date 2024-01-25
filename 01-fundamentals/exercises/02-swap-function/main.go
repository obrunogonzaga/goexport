package main

import (
	"os"
	"strconv"
)

// Create a `swap` function that takes two integer pointers and swaps their values.
func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		println("Please provide two numbers.")
		return
	}

	arg1 := args[0]
	arg2 := args[1]

	num1, err1 := strconv.Atoi(arg1)
	num2, err2 := strconv.Atoi(arg2)
	if err1 != nil || err2 != nil {
		println("Please provide two numbers.")
		return
	}

	println("Before swap:", num1, num2)
	swap(&num1, &num2)
	println("After swap:", num1, num2)

}

func swap(num1 *int, num2 *int) {
	temp := *num1
	*num1 = *num2
	*num2 = temp
}
