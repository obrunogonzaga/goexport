package main

import "net/http"

func main() {
	http.HandleFunc("/busca-cep", BuscaCepHandler)
	http.ListenAndServe(":8080", nil)
}

func BuscaCepHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello World!\n"))
	if err != nil {
		return
	}
}
