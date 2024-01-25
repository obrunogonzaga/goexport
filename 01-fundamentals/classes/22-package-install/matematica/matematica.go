package matematica

func Soma[T int | float64](x, y T) T {
	return x + y
}

var A = 10

type Carro struct {
	Marca string
}

func (c Carro) Andar() string {
	return "Andando..."
}
