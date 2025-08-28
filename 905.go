package main
import (
	b "bufio"
	. "fmt"
	. "os"
)
type T map[byte]byte
func main() {
	s := b.NewScanner(Stdin)
	var (
		e []string
		i = 0
		p = T{}
		x = -1
		a = "the quick brown fox jumps over the lazy dog"
		P = Println
	)

	Scanln(&i)
	for 0 < i {
		s.Scan()
		e = append(e, s.Text())
		i--
	}

	for i, l := range e {
		if len(l) == len(a) {
			o := T{}
			u := T{}
			v := 1
			j := 0
			for j < len(l) {
				h := l[j]
				z := a[j]
				if h != 32 {
					if z == 32 {
						v = 0
						break
					}
					w, t := o[h]
					if t {
						if w != z {
							v = 0
							break
						}
					} else {
						if u[z] > 0 {
							v = 0
							break
						}
						o[h] = z
						u[z] = 1
					}
				} else if z != 32 {
					v = 0
				}
				j++
			}
			if v > 0 {
				p = o
				x = i
				break
			}
		}
	}

	if x < 0 {
		P("No solution")
		return
	}

	for _, y := range e {
		j := make([]byte, len(y))
		i = 0
		for i < len(y) {
			j[i] = 32
			if y[i] != 32 {
				q, o := p[y[i]]
				j[i] = 63
				if o {
					j[i] = q
				}
			}
			i++
		}
		P(string(j))
	}
}