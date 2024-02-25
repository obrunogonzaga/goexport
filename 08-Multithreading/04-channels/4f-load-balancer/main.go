package main

import (
	"fmt"
	"time"
)

func worker(workedId int, c chan int) {
	for x := range c {
		fmt.Printf("Worker %d received %d\n", workedId, x)
		time.Sleep(time.Second)
	}
}

func main() {
	c := make(chan int)
	QtdWorkers := 1000000

	for i := 0; i < QtdWorkers; i++ {
		go worker(i, c)
	}

	for i := 0; i < 10000000; i++ {
		c <- i
	}
}
