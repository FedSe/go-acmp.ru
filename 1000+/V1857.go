package main
import (
	b "bufio"
	. "fmt"
	. "os"
)
func main() {
	var n, m, d, x, y, i int
	r := b.NewReader(Stdin)

	Scan(&n, &m, &d)
	for i < d {
		Fscan(r, &x, &y)
		if x < n {
			n = x
		}
		if y < m {
			m = y
		}
		i++
	}

	Print(n*m, d)
}