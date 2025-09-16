package main
import . "fmt"
func main() {
	var (
		a       [2e3]int
		p, q, n int
		P       = Println
	)

	Scan(&p, &q)
	for p > 0 {
		n++
		p *= n
		a[n] = p / q
		p %= q
	}

	P(n)
	q = 2
	for q < n {
		P(a[q])
		q++
	}
	P(a[n])
}