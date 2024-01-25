package main

import (
	"fmt"
	"github.com/obrunogonzaga/pos-go-expert/matematica"
)

func main() {
	soma := matematica.Soma(3, 2)
	carro := matematica.Carro{}
	carro.Marca = "Fiat"
	fmt.Println("Resultado: ", soma)
	fmt.Println("A: ", matematica.A)
	fmt.Println("Carro: ", carro.Marca)
	fmt.Println("Carro: ", carro.Andar())
}
