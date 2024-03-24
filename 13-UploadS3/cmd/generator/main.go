package main

import (
	"fmt"
	"os"
)

func main() {
	i := 0
	for {
		f, err := os.Create(fmt.Sprintf("./tmp/file-%d.txt", i))
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		f.Write([]byte("Hello, World!"))
		i++
	}
}
