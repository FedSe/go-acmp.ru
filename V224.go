package main
import (
	b "bufio"
	. "fmt"
	. "os"
	. "sort"
	. "strconv"
)
func main() {
	var (
		a    [1e6]int
		n, i int
		s    = b.NewScanner(Stdin)
	)
	s.Split(b.ScanWords)

	Scan(&n)
	for s.Scan() {
		a[i], _ = Atoi(s.Text())
		i++
	}

	Ints(a[:n])
	i = a[n-1]
	f := a[0] * a[1] * i
	g := i * a[n-2] * a[n-3]

	if f > g {
		g = f
	}
	Print(g)
}