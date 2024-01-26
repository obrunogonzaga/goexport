package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int
	Saldo  int
}

type ContaTag struct {
	Numero int `json:"number"`
	Saldo  int `json:"balance"`
}

func main() {
	// Transformando um objeto em JSON
	conta := Conta{Numero: 123, Saldo: 1000}
	res, err := json.Marshal(conta)
	if err != nil {
		panic(err)
	}
	println(string(res))

	// Transformando um objeto em JSON - Utilizando o Encode
	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		panic(err)
	}

	// Transformando um JSON em objeto - Unmarshal
	jsonPuro := []byte(`{"Numero":123,"Saldo":1000}`)
	var contaX Conta
	err = json.Unmarshal(jsonPuro, &contaX)
	if err != nil {
		panic(err)
	}
	println(contaX.Saldo)

	// Transformando um JSON em objeto - Unmarshal - JSON diferente do Obj
	jsonPuroTag := []byte(`{"number":123,"balance":1000}`)
	var contaXTag ContaTag
	err = json.Unmarshal(jsonPuroTag, &contaXTag)
	if err != nil {
		panic(err)
	}
	println(contaXTag.Saldo)
}
