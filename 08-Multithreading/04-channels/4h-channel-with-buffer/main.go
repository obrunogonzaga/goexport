package main

func main() {
	ch := make(chan string, 2) //Empty channel - buffer 2
	ch <- "Hello,"
	ch <- "World!"

	println(<-ch)
	println(<-ch)
}
