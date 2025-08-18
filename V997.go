package main
import (
	b "bufio"
	. "fmt"
	. "os"
	. "strings"
)
func main() {
	var (
		N, M int
		y    = map[any]int{}
		z    = map[any]int{}
		w    []string
		s    = b.NewScanner(Stdin)
		a    = "Everything is going to be OK."
		o    = ""
		q    = o
		L    = ToLower
	)

	Scanln(&N, &M)
	for 0 < N {
		Scanln(&o)
		z[L(o)] = 1
		w = append(w, L(o))
		N--
	}

	for 0 < M {
		s.Scan()
		for _, v := range " " + s.Text() {
			if v > 64 {
				if v < 91 {
					v += 32
				}
				q += string(v)
			} else {
				if len(q) > 0 {
					y[q] = 1
					q = ""
				}
			}
		}
		M--
	}

	if len(q) > 0 {
		y[q] = 1
	}

	for _, v := range w {
		if y[v] < 1 {
			a = "The usage of the vocabulary is not perfect."
		}
	}

	for v := range y {
		if z[v] < 1 {
			a = "Some words from the text are unknown."
		}
	}

	Print(a)
}