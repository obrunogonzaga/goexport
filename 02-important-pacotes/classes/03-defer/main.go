package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	request, err := http.Get("http://www.coritiba.com.br")
	if err != nil {
		panic(err)
	}
	defer request.Body.Close()

	result, err := io.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}
	println(string(result))

	fmt.Println("First")
	defer fmt.Println("Second")
	fmt.Println("Third")
}
