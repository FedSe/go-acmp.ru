package main
import . "fmt"
func main() {
	var (
		n, p int
		S    [1e5]int
		k    = 2
	)

	Scan(&n)
	for k*k < n*2 {
		l := k * k
		for l < 2*n {
			S[l] = 1
			l += k
		}
		k++
	}

	k = n + 1
	for k < n*2 {
		if S[k] < 1 {
			p++
		}
		k++
	}

	Print(p)
}