package main
import . "fmt"
func main() {
	var n, s, i int
	S := Scan

	S(&n)
	a := make([]int, n)
	for i < n {
		S(&a[i])
		i++
	}

	S(&n)
	for _, d := range a {
		if d > n {
			d = n
		}
		s += d
	}
	Print(s)
}