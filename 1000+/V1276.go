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
		S             = Fscan
	)

	S(s, &N)
	for j < N {
		S(s, &a[j])
		j++
	}

	for l < N {
		o++
		j = a[l]
		for l < N && a[l] == j {
			l++
		}
	}

	Println(o)
	for i < N {
		j = a[i]
		l = 0
		for i < N && a[i] == j {
			i++
			l++
		}
		Printf("[%d %d] ", j, l)
	}
}