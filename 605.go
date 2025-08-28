package main
import . "fmt"
type H struct {
	s int
	n string
}
func main() {
	var (
		o, h []H
		n, i int
		y    [][]string
		S    = Sprint
		P    = Print
	)

	Scan(&n)
	for i < 20 {
		i++
		o = append(o, H{i, S(i)})
		o = append(o, H{2 * i, "D" + S(i)})
		o = append(o, H{3 * i, "T" + S(i)})
		h = append(h, H{2 * i, "D" + S(i)})
	}
	o = append(o, H{25, "25"})
	o = append(o, H{50, "Bull"})
	h = append(h, H{50, "Bull"})

	for _, v := range h {
		if v.s <= n {
			r := n - v.s
			if r == 0 {
				y = append(y, []string{v.n})
			} else {
				for _, a := range o {
					if a.s == r {
						y = append(y, []string{a.n, v.n})
					}
				}
				for _, a := range o {
					for _, b := range o {
						if a.s+b.s == r {
							y = append(y, []string{a.n, b.n, v.n})
						}
					}
				}
			}
		}
	}

	P(len(y), `
`)
	for _, r := range y {
		for _, s := range r {
			P(s, " ")
		}
		P(`
`)
	}
}