package main
import . "fmt"
func main() {
	var (
		n, p int
		S [100000]int
	)
	Scan(&n)

	k := 2
	for k*k < n*2 {
		for
		l := k*k
		l < 2*n
		{
			S[l] = 1
		l += k
		}
	k++
	}

	k = n+1
	for k < n*2 {
		if S[k] < 1 {
			p++
		}
	k++
	}
	Print(p)
}