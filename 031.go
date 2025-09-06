package main
import . "fmt"
func main() {
	var (
		n, k, s, i int
		c          = 1
		t          = 1
	)

	Scan(&n, &k)
	j := k
	for j < n {
		j++
		c *= j
	}
	j = 1
	for i <= n-k {
		s += j * c / t
		j = -j
		i++
		t *= i
	}

	Print(s)
}