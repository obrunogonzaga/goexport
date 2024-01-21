package main

import (
	"fmt"
)

func main() {

	total := func() int {
		return soma(1, 3, 5, 7, 11, 18) * 2
	}()
	fmt.Println(total)
}

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
