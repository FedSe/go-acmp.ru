package main
import (
	. "fmt"
	b "bufio"
	. "os"
)
func main() {
	var (
		a [1440]int
		p, f, u, m, i int
		s = b.NewScanner(Stdin)
	)
	Scan(&p)

	for s.Scan() {
		t := s.Text()
		Sscanf(t, "%d:%d %d:%d", &p, &f, &u, &m)
		f += p*60
		for f <= m+u*60 {
			a[f]++
			if a[f] > i { i = a[f] }
		f++
		}
	}

	Print(i)
}