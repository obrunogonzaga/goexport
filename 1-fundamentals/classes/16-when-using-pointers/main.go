package main

func soma(a, b *int) int {
	*a = 50
	return *a + *b
}

func main() {
	a, b := 1, 2
	println(soma(&a, &b))
	println(a, b)
}
