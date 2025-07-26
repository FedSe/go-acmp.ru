package main
import . "fmt"
func main() {
	var n, k, s int
	Scan(&n, &k)
	p := 1

	for n > 0 {
		s += n % k
		p *= n % k
		n /= k
	}

	Print(p - s)
}