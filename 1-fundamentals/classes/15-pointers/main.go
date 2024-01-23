package main

import "fmt"

func main() {
	numero := 10
	fmt.Printf("Endereço de memória: %p\n", &numero)
	fmt.Printf("Valor da variável: %d\n", numero)

	var ponteiro *int = &numero
	fmt.Printf("Endereço de memória: %p\n", ponteiro)
	fmt.Printf("Valor da variável: %d\n", *ponteiro)

	*ponteiro = 20
	fmt.Printf("Endereço de memória: %p\n", &numero)
	fmt.Printf("Valor da variável: %d\n", numero)


	numero2 := &numero
	fmt.Printf("Endereço de memória: %p\n", numero2)
	fmt.Printf("Valor da variável: %d\n", *numero2)

	*numero2 = 30
	fmt.Printf("Endereço de memória: %p\n", &numero)
	fmt.Printf("Valor da variável: %d\n", numero)

}
