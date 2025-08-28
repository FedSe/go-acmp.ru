package main
import . "fmt"
func main() {
	n := 0
	i := 2

	Scan(&n)
	r := n
	for i*i < n {
		if n%i < 1 {
			r -= r / i
		}
		for n%i < 1 {
			n /= i
		}
		i++
	}

	Print(r - r/n)
}