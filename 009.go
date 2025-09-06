package main
import . "fmt"
func main() {
	var (
		a, s, p, k, m, L, X, v int
		N                      [100]int
	)

	Scan(&a)
	for p < a {
		Scan(&v)
		N[p] = v
		if p < 1 {
			X = v
			L = X
		}
		if v > 0 {
			s += v
		}
		if v > X {
			X = v
			k = p
		}
		if v < L {
			L = v
			m = p
		}
		p++
	}

	p = 1
	if k > m {
		k, m = m, k
	}

	for k < m-1 {
		k++
		p *= N[k]
	}

	Print(s, p)
}