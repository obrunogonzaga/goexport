package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
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
	http.HandleFunc("/busca-cep", BuscaCepHandler)
	http.ListenAndServe(":8080", nil)
}

func BuscaCepHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/busca-cep" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		http.Error(w, "cep is required.", http.StatusBadRequest)
		return
	}

	cep, err := BuscaCep(cepParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(cep)
	if err != nil {
		panic(err)
	}
}

func BuscaCep(cep string) (*Address, error) {
	res, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	var address Address
	err = json.Unmarshal(body, &address)
	return &address, nil
}
