package main
import (
	b "bufio"
	. "fmt"
	. "os"
)
func main() {
	var (
		c, v    [1e3]int
		n, i, j int
		s       = b.NewReader(Stdin)
		S       = Fscan
	)

	S(s, &n)
	for j < n {
		S(s, &v[j])
		j++
	}

	for i < n {
		j = v[i]
		c[2] += c[1] * j
		c[1] += c[0] * j
		c[0] += j
		i++
	}

	Print(c[2])
}