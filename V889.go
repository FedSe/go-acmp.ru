package main
import (
	b "bufio"
	. "fmt"
	. "os"
	. "sort"
)
func main() {
	s := 0
	i := 0
	r := b.NewReader(Stdin)

	Fscan(r, &s, &i)

	h := make([][2]int, i)

	for i > 0 {
		i--
		Fscan(r, &h[i][0], &h[i][1])
	}

	Slice(h, func(i, j int) bool {
		return h[i][1] > h[j][1]
	})

	for _, h := range h {
		if h[0] == s {
			s++
		} else if h[0]+1 == s {
			s--
		}
	}

	Print(s)
}