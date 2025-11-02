package main
import (
	b "bufio"
	. "fmt"
	. "os"
	. "strings"
)
type S = string
func main() {
	s := b.NewScanner(Stdin)
	s.Scan()

	d := Fields(s.Text())
	m := len(d)
	i := 0
	e := make([]S, m)
	w := make([]S, m)
	for i, v := range d {
		n := len(v)
		t := make([]byte, n)
		n /= 2
		l := v[:n]
		v = v[n:]
		j := 0
		for j < len(v) {
			t[2*j] = v[j]
			j++
		}
		j = 0
		n = len(l)
		for j < n {
			t[2*j+1] = l[n-1-j]
			j++
		}
		w[i] = S(t)
	}

	m /= 2
	l := w[:m]
	w = w[m:]
	for i < len(w) {
		e[2*i] = w[i]
		i++
	}
	i = 0
	m = len(l)
	for i < m {
		e[2*i+1] = l[m-1-i]
		i++
	}

	Print(Join(e, " "))
}