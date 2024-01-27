package main

import "net/http"

func main() {
	fileSerever := http.FileServer(http.Dir("./public")) // FileServer retorna um manipulador que serve arquivos no sistema de arquivos do sistema operacional
	mux := http.NewServeMux()                            // NewServeMux cria um novo ServeMux, que é um roteador HTTP simples
	mux.Handle("/", fileSerever)                         // Handle registra o manipulador para o padrão de URL especificado
	http.ListenAndServe(":8080", mux)                    // ListenAndServe ouve as requisições HTTP na porta 8080
}
