package main
import . "fmt"
func main() {
	var (
		N, K, i int
		f       = 1
		c       = 1
	)

	Scan(&N, &K)
	for i < K {
		c *= N - i
		i++
		c /= i
		f *= i
	}
	if K > N {
		c = 0
	}

	Print(c * c * f)
}