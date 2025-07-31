package main
import (
	. "fmt"
	. "sort"
)

type S struct {
	e, d int
}

func main() {
	var (
		f    []S
		n, d int
	)

	Scan(&n)
	for d <= n {
		j := 1
		for j < d {
			a := j
			b := d
			for b > 0 {
				a, b = b, a%b
			}
			if a == 1 {
				f = append(f, S{j, d})
			}
			j++
		}
		d++
	}

	Slice(f, func(i, j int) bool {
		return f[i].e*f[j].d < f[j].e*f[i].d
	})

	for _, v := range f {
		Print(v.e, "/", v.d, `
`)
	}
}