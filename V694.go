package main
import . "fmt"
func main() {
	var (
		a, b, n, i, k int
		s = "NO"
		d [31]int
	)
	Scan(&n)

	for i < n {
		Scan(&a, &b)
		for
		j := a
		j <= b
		{
			d[j-1]++
		j++
		}
	i++
	}

	for k < 31 {
		if d[k] == n { s = "YES"}
	k++
	}

	Print(s)
}