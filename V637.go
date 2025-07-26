package main
import . "fmt"
func main() {
	var n, s, i int

	Scan(&n)
	a := make([]int, n)
	for i < n {
		Scan(&a[i])
	i++
	}

	Scan(&n)
	for _ , d := range a {
		if d > n { d = n }
		s += d
	}
	Print(s)
}