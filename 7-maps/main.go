package main

import "fmt"

func main() {
	salarios := map[string]float64{"Bruno": 2000.00, "Alzi": 4000.00, "Gabriela": 3000.00}
	fmt.Println(salarios)
	fmt.Println(salarios["Bruno"])
	delete(salarios, "Bruno")
	salarios["Gonzaga"] = 10000
	fmt.Println(salarios["Gonzaga"])
	salario := make(map[string]float64)
	//sal := map[string]int{} outra opção
	salario["Bruno"] = 10000	

	for nome, salario := range salarios {
		fmt.Printf("%s tem salario de R$ %.2f\n", nome, salario)
	}

	for _, salario := range salarios {
		fmt.Printf("O salario de R$ %.2f\n", salario)
	}

}
