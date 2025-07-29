package main
import . "fmt"
func main() {
	var n, k, i int
	r := 1

	Scan(&n, &k)
	for i < k {
		i++
		r = r * (n - i + 1) / i
	}

	Print(r)
}