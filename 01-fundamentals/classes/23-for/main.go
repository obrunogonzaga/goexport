package main

func main() {
	// for init; condition; post {}
	for i := 0; i < 10; i++ {
		println(i)
	}

	// for condition {}
	j := 0
	for j < 10 {
		println(j)
		j++
	}

	// for {} - loop infinito
	k := 0
	for {
		println(k)
		k++
		if k == 10 {
			break
		}
	}

	// for range
	s := []int{1, 2, 3}
	for i, v := range s {
		println(i, v)
	}
}
