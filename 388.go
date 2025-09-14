package main
import (
	b "bufio"
	. "fmt"
	. "os"
)
func main() {
	var (
		n, m, c, x, i, j, a int
		R, C, D             [6e5]int
		s                   = b.NewReader(Stdin)
	)

	Scan(&n, &m)
	for x < m {
		C[x] = -2e3
		x++
	}
	w := n * m
	for j < w {
		Fscan(s, &D[j])
		j++
	}

	for i < n {
		R[i] = 2e3
		j = 0
		for j < m {
			a = D[i*m+j]
			if a < R[i] {
				R[i] = a
			}
			if a > C[j] {
				C[j] = a
			}
			j++
		}
		i++
	}

	for 0 < n {
		n--
		j = 0
		for j < m {
			if R[n] == C[j] {
				c++
			}
			j++
		}
	}

	Print(c)
}