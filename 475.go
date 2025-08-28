package main
import (
	b "bufio"
	. "fmt"
	. "os"
	. "sort"
)
func main() {
	var (
		a [1e5]int
		n = 0
		r = b.NewScanner(Stdin)
		s = "Yes"
	)
	r.Split(b.ScanWords)

	for r.Scan() {
		Sscan(r.Text(), &a[n])
		n++
	}

	Ints(a[:n])
	d := a[1] - a[0]

	for 1 < n {
		n--
		if a[n]-a[n-1] != d {
			s = "No"
		}
	}

	Print(s)
}