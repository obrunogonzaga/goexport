package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	bruno := Cliente{
		Nome:  "Bruno",
		Idade: 37,
		Ativo: true,
	}
	bruno.Ativo = false
	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", bruno.Nome, bruno.Idade, bruno.Ativo)
}
