package main
import (
	. "fmt"
	. "sort"
)
func main() {
	var (
		n, i, j int
		c, d    []string
		S       = Scan
	)

	S(&n)
	a := make([]string, n)
	for i < n {
		S(&a[i])
		i++
	}

	S(&n)
	b := make([]string, n)
	for j < n {
		S(&b[j])
		j++
	}

	for _, v := range a {
		for _, w := range b {
			if v == w {
				c = append(c, v)
			}
		}
	}

	for _, v := range b {
		n = 1
		for _, w := range c {
			if v == w {
				n = 0
			}
		}
		if n > 0 {
			d = append(d, v)
		}
	}

	for _, v := range []string{`Friends: `, `
Mutual Friends: `, `
Also Friend of: `} {
		Strings(a)
		for i, s := range a {
			if i < len(a)-1 {
				s += ", "
			}
			v += s
		}
		Print(v)
		a = c
		c = d
	}
}