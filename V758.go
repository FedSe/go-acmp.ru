package main
import (
	b "bufio"
	. "fmt"
	. "os"
	. "sort"
)
func main() {
	var (
		a, n, i int
		t       []int
		s       = b.NewReader(Stdin)
		S = Fscan
	)

	S(s, &n)
	for n > 0 {
		S(s, &i)
		t = append(t, i)
		n--
	}
	Ints(t)

	for len(t) > 3 {
		i = t[1] * 2
		n = len(t) - 1
		a += t[n] + t[0]
		if t[n-1]+t[0] > i {
			a += i
			n--
		}
		t = t[:n]
	}

	a += t[0]
	if len(t) > 2 {
		a += t[1] + t[2]
	}
	if len(t) == 2 {
		a += t[1] - t[0]
	}

	Print(a)
}