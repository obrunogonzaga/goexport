package main

import "net/http"

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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Hello World!\n"))
	if err != nil {
		return
	}
}
