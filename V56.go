package main
import (
	. "fmt"
	. "sort"
)
func main() {
	var (
		n, i, j int
		c, d []string
	)

	Scan(&n)
	a := make([]string, n)
	for i < n {
		Scan(&a[i])
	i++
	}

	Scan(&n)
	b := make([]string, n)
	for j < n {
		Scan(&b[j])
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

	for _ , i := range [3]string{"Friends: ", "Mutual Friends: ", "Also Friend of: "} {
		Print("\n", i)
		Strings(a)
		for i, s := range a {
			if i != len(a)-1 {
				s += ", "
			}
			Print(s)
		}
	a = c
	c = d
	}
}