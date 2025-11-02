package main
import . "fmt"
func main() {
	var (
		z    [1e3]int
		n, i int
	)

	Scan(&n)
	if n == 0 {
		i = 1
	}
	for n != 0 {
		z[i] = (n%-3 + 3) % 3
		n -= z[i]
		n /= -3
		i++
	}

	for i > 0 {
		i--
		Print(z[i])
	}
}