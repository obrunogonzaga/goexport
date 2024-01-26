package main

import (
	"io"
	"net/http"
)

func main() {
	request, err := http.Get("http://www.coritiba.com.br")
	if err != nil {
		panic(err)
	}
	result, err := io.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}
	println(string(result))
	request.Body.Close()
}
