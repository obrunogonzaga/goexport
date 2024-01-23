package main

import "fmt"

// Define a struct representing a `Person` with `Name` and `Age` fields. Write a function that takes a pointer to a `Person` and modifies its `Age`.
type Person struct {
	Name string
	Age  int
}

func modifyAge(p *Person, age int) {
	p.Age = age
}

func main() {
	bruno := Person{"Bruno", 37}
	fmt.Printf("Age of %s is %d\n", bruno.Name, bruno.Age)
	modifyAge(&bruno, 38)
	fmt.Printf("Age of %s is %d\n", bruno.Name, bruno.Age)
}
