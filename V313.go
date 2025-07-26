package main
import (
	. "fmt"
	b "bufio"
	. "os"
)
func main() {
	var (
		t [202]int
		z, n, a, m int
		s = b.NewScanner(Stdin)
	)
	s.Split(b.ScanWords)

	Scanln(&n)
	for s.Scan() {
	z++
		Sscan(s.Text(), &a)
		f := &t[a+101]
		b := z - *f
		if b > t[a] && *f > 0 {
			t[a] = b
		}
		*f = z
	}

	z = 1
	for z < 101 {
		if m < t[z] {
			m = t[z]
		}
	z++
	}

	Print(m)
}