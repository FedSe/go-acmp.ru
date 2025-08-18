package main
import (
	b "bufio"
	. "fmt"
	. "os"
	. "sort"
)
func main() {
	var (
		A          [1e5]int
		n, m, t, i int
		r          = b.NewReader(Stdin)
	)

	Scan(&n)
	for t < n {
		Fscan(r, &A[t])
		t++
	}

	Ints(A[:n])
	for i < n {
		t = A[i] + n - 1 - i
		if t > m {
			m = t
		}
		i++
	}

	Print(m)
}