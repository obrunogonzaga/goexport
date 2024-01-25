package main

func main() {
	a := 10
	b := 100
	c := 1000

	if a < b && b < c {
		println("both conditions are true")
	}

	if a < b || b < c {
		println("at least one condition is true")
	}

	if a == b {
		println("a is equal to b")
	} else {
		println("a is not equal to b")
	}

	if a != b {
		println("a is not equal to b")
	}

	if a < b {
		println("a is less than b")
	}

	if a > b {
		println("a is greater than b")
	}

	if a <= b {
		println("a is less than or equal to b")
	}

	if b >= a {
		println("b is greater than or equal to a")
	}

	switch a {
	case 10:
		println("a is equal to 0")
	case 100:
		println("a is equal to 1")
	case 1000:
		println("a is equal to 2")
	default:
		println("a is not equal to 0, 1 or 2")
	}

}
