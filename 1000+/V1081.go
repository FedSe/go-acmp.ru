package main
import (
	b "bufio"
	. "fmt"
	. "os"
)
func main() {
	var (
		n, m, i, l int
		r          = b.NewReader(Stdin)
		S          = Fscan
	)

	S(r, &n)
	p := make([]int, n+1)

	for i < n {
		i++
		S(r, &p[i])
	}

	for l < n {
		l++
		p[l] += p[l-1]
	}

	S(r, &m)
	for 0 < m {
		S(r, &n, &l)
		Println(p[l] - p[n-1])
		m--
	}
}