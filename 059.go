package main
import . "fmt"
func main() {
	var n, k, s int
	p := 1

	Scan(&n, &k)
	for n > 0 {
		s += n % k
		p *= n % k
		n /= k
	}

	Print(p - s)
}