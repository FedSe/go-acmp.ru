package main
import . "fmt"
func main() {
	var n, k, s, i int

	Scan(&n, &k)
	for i < n {
		i++
		s += k
		s %= i
	}

	Print(s + 1)
}