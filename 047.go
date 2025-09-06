package main
import . "fmt"
func main() {
	var a, b, n, i int

	Scan(&n)
	for i <= n {
		i++
		s := 0
		x := i
		for x > 0 {
			s += x % 10
			x /= 10
		}
		if n%i < 1 && s > b {
			a = i
			b = s
		}
	}

	Print(a)
}