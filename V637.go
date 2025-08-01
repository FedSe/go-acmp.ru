package main
import . "fmt"
func main() {
	var (
		n, s int
		a    [1e4]int
		S    = Scan
	)

	S(&n)
	for 0 < n {
		n--
		S(&a[n])
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