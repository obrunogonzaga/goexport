package main

import "fmt"

const a = "Hello, World!"
var	(
		b bool = true
		c int
		d string
		e float64
		g ID = 1
)

type ID int

func main()  {
	f := "F"	//Cria variavel inferindo o valor, usa somente a primeira vez
	f = "f"
	println(a)
	println(b)
	println(c)
	println(d)
	println(e)
	println(f)
	println(g)
	fmt.Printf("%T\n", g) 
}