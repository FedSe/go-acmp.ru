package main
import . "fmt"
func main() {
	var n, m, u, r, v int

	Scan(&n, &m)
	if m < 1 || n == m {
		u = 1
		goto A
	}
	if m == 1 {
		u = n
		goto A
	}
	if n < 1 || n < m {
		goto A
	}

	for m+m*r-r <= n {
		v++
		r++
	}

	m--
	u = v*n - v*(v+1)/2*m
A:
	Print(u)
}