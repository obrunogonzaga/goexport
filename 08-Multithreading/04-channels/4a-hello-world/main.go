package main

import "fmt"

func main() {
	canal := make(chan string) //Empty channel

	go func() {
		canal <- "Olá, canal!" //Full channel
	}()

	msg := <-canal //channel is emptying
	fmt.Println(msg)
}
