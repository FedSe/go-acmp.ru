package main
import . "fmt"
func main() {
	n := 0
	Scan(&n)

	a := n / 107
	n %= 107
	b := a*7
	a *= 100

	if n > 7 {
		a += n-7
		b += 7
	}

	Print(a, b)
}