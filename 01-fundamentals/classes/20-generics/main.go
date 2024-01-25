package main

type MyNumber int

type Number interface {
	~int | float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Compara[T comparable](a, b T) bool {
	return a == b
}

func main() {
	m := map[string]int{"Bruno": 1000, "Alzi": 2000, "Gabriela": 3000}
	println(Soma(m))

	m2 := map[string]float64{"Bruno": 1000.10, "Alzi": 2000.20, "Gabriela": 3000.30}
	println(Soma(m2))

	m3 := map[string]MyNumber{"Bruno": 1000, "Alzi": 2000, "Gabriela": 3000}
	println(Soma(m3))

	println(Compara(10, 10.0))
}
