package main
import (
	b "bufio"
	. "fmt"
	. "os"
)
func main() {
	var (
		a             [1440]int
		p, f, u, m, i int
		s             = b.NewReader(Stdin)
	)

	Fscanln(s, &p)
	l := p
	for l > 0 {
		Fscanf(s, `%d:%d %d:%d
`, &p, &f, &u, &m)
		f += p * 60
		for f <= m+u*60 {
			a[f]++
			if a[f] > i {
				i = a[f]
			}
			f++
		}
		l--
	}

	Print(i)
}