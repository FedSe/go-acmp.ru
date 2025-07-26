package main
import (
	. "fmt"
	b "bufio"
	. "os"
	. "strconv"
)
func main() {
	var (
		n, m, c, x, i, f, j, a int
		R, C, D [564001]int
		s = b.NewScanner(Stdin)
	)
	s.Split(b.ScanWords)
	Scan(&n, &m)

	for x < m {
		C[x] = -1001
	x++
	}

	for s.Scan() {
		D[j], _ = Atoi(s.Text())
	j++
	}

	for i < n {
		R[i] = 1001
		for
		j = 0
		j < m
		{
			a = D[i*m+j]
			if a < R[i] { R[i] = a }
			if a > C[j] { C[j] = a }
		j++
		}
	i++
	}

	for f < n {
		for
		j = 0
		j < m
		{
			if R[f] == C[j] { c++ }
		j++
		}
	f++
	}

	Print(c)
}