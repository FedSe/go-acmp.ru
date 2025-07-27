package main
import . "fmt"

func P(a int) {
	Println(a)
}

func main() {
	n := 0
	Scan(&n)

	if n > 8 {
		P(n - 8)
	}

	if (n+7)%8 > 0 {
		P(n - 1)
	}

	if n%8 > 0 {
		P(n + 1)
	}

	if n < 57 {
		P(n + 8)
	}
}