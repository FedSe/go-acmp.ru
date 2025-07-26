package main
import (
	. "fmt"
	. "sort"
	b "bufio"
	. "os"
)
func main() {
	var (
		a, n int
		t [] int
		s = b.NewScanner(Stdin)
	)
	s.Split(b.ScanWords)
	Scan(&n)

	for s.Scan() {
		Sscan(s.Text(), &n)
		t = append(t, n)
	}
	Ints(t)

	for len(t) > 3 {
		i := t[1]*2
		n = len(t)-1
		a += t[n] + t[0]
		if t[n-1] + t[0] > i {
			a += i
			n--
		}
		t = t[:n]
	}

	a += t[0]
	if len(t) > 2 { a += t[1] + t[2] }
	if len(t) == 2 { a += t[1] - t[0] }

	Print(a)
}