package main
import (
	b "bufio"
	."fmt"
	. "os"
	. "sort"
	. "strconv"
)
func main() {
	n := 0
	i := 0
	s := b.NewScanner(Stdin)
	s.Split(b.ScanWords)

	Scan(&n)
	a := make([]int, n)
	for s.Scan() {
		a[i], _ = Atoi(s.Text())
	i++
	}

	Ints(a)
	i = a[n-1]
	f := a[0] * a[1] * i
	g := i * a[n-2] * a[n-3]

	if f > g { g = f }
	Print(g)
}