package main
import (
	b "bufio"
	. "fmt"
	. "os"
)
func main() {
	var (
		s             = b.NewReader(Stdin)
		N, j, o, l, i int
		a             [1e6]int
	)

	Scan(&N)
	for j < N {
		Fscan(s, &a[j])
		j++
	}

	for l < N {
		o++
		j = a[l]
		for l < N && a[l] == j {
			l++
		}
	}

	Print(o)
	for i < N {
		j = a[i]
		l = 0
		for i < N && a[i] == j {
			i++
			l++
		}
		Printf(" [%d %d]", j, l)
	}
}