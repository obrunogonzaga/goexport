package main

import (
	"errors"
	"fmt"
)

func main() {
	valor, err := soma(5, 4)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(valor)
}

func soma(num1, num2 int) (int, error) {
	if num1 > 10 {
		return 0, errors.New("numero maior que 10")
	}
	return num1 + num2, nil
}
