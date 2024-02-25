package main

import "fmt"

func main() {
	hello := make(chan string)
	go recebe("Olá, canal!", hello)
	ler(hello)

}

// Received only
// <- do lado direito do canal, indica que o canal só pode receber valores (encher canal)
func recebe(nome string, hello chan<- string) {
	hello <- nome
}

// Send only
// <- do lado esquerdo do canal, indica que o canal só pode enviar valores (esvaziar canal)
func ler(hello <-chan string) {
	fmt.Println(<-hello)
}
