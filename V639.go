package main
import (
	. "fmt"
	. "sort"
)
func main() {
	var (
		n, u, i int
		c       []float64
		e       []string
		o       []int
		r       = .0
	)
	Scan(&n)

	for 0 < n {
		Scan(&u)
		for 0 < u {
			Scan(&r)
			c = append(c, r)

			a := ""
			Scan(&a)
			e = append(e, a)
			o = append(o, len(o))
			u--
		}
		n--
	}

	Slice(o, func(i, j int) bool {
		return c[o[i]] > c[o[j]]
	})

	u = len(o)
	Println(u)
	for i < u {
		Printf(`%.2f %s
`, c[o[i]], e[o[i]])
		i++
	}
}