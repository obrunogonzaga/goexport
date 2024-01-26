package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Address struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, cep := range os.Args[1:] {
		// Get URL with CEP [via CEP]
		req, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
		if err != nil {
			_, err := fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			if err != nil {
				panic(err)
			}
		}
		defer req.Body.Close()

		// Read response body
		res, err := io.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}

		// Parse response body
		var address Address
		err = json.Unmarshal(res, &address)
		if err != nil {
			panic(err)
		}

		// Write in file
		file, err := os.Create("address.txt")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		_, err = file.WriteString(fmt.Sprintf("CEP: %s, Logradouro: %s, Complemento: %s, Bairro: %s, Localidade: %s, UF: %s, DDD: %s",
			address.Cep, address.Logradouro, address.Complemento, address.Bairro, address.Localidade, address.Uf, address.Ddd))
		if err != nil {
			panic(err)
		}
		println("File created successfully")
	}
}
