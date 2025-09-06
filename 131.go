package main
import . "fmt"
func main() {
	var i, n, v, a, s int
	b := -1

	Scan(&n)
	for i < n {
		i++
		Scan(&v, &s)
		if s > 0 && v > a {
			a = v
			b = i
		}
	}

	Print(b)
}