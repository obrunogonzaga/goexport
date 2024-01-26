package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// Gravando um arquivo
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	file.WriteString("Teste de escrita em arquivo")

	file.Close()

	// Gravando bytes em um arquivo
	file, err = os.Create("test2.txt")
	if err != nil {
		panic(err)
	}
	file.Write([]byte("Teste de escrita em arquivo"))
	if err != nil {
		panic(err)
	}
	file.Close()

	// Lendo um arquivo
	fileReader, err := os.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(fileReader))

	// Ler arquivo pouco a pouco
	arquivo, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(arquivo)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Println(line)
	}
	arquivo.Close()

	// Ler arquivo pouco a pouco com buffer
	arquivo2, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	reader2 := bufio.NewReader(arquivo2)
	buffer := make([]byte, 10)
	for {
		n, err := reader2.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}
	arquivo2.Close()

	// Removendo um arquivo
	err = os.Remove("test.txt")
	if err != nil {
		panic(err)
	}
	err = os.Remove("test2.txt")
	if err != nil {
		panic(err)
	}
}
