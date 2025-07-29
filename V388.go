package main
import (
	b "bufio"
	. "fmt"
	. "os"
)
func main() {
	var (
		n, m, c, x, i, f, j, a int
		R, C, D                [564001]int
		s                      = b.NewReader(Stdin)
		S                      = Fscan
	)

	S(s, &n, &m)
	for x < m {
		C[x] = -1001
		x++
	}
	w := n * m
	for j < w {
		S(s, &D[j])
		j++
	}

	for i < n {
		R[i] = 1001
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

	for f < n {
		j = 0
		for j < m {
			if R[f] == C[j] {
				c++
			}
			j++
		}
		f++
	}

	Print(c)
}