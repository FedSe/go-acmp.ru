package main
import (
	b "bufio"
	. "fmt"
	. "os"
	. "strings"
)
func main() {
	s := b.NewScanner(Stdin)
	w := []string{}
	for s.Scan() {
		w = append(w, s.Text())
	}

	p := Fields(w[0])
	u := p[len(p)-3]
	for i, z := range w[1 : len(w)-1] {
		k := "Fedya"
		if i%2 > 0 {
			k = u
		}
		x := z[Index(z, ": ")+2:]
		w := len(x) - 1
		l := x[w]
		q := x + ","
		if l == 46 {
			q = x[:w] + ","
		}
		if l == 33 || l == 63 {
			q = x[:w] + string(l)
		}
		Printf(`"%s" --- skazal %s.
`, q, k)
	}
}