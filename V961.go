package main
import (
	. "fmt"
	. "sort"
)
func main() {
	var (
		n, m, j, i, l int
		u, w []string
		a [100]string
		e = ""
	)

	Scan(&n, &m)

	for j < n {
		Scan(&a[j])
	j++
	}

	for i < n {
	Scan(&e)
		for
		j = 0
		j < m
		{
			c := a[i][j]
			v := string(c)
			if c != e[j] && c != 46 {
				if c < 91 {
					u = append(u, v)
				} else {
					w = append(w, v)
				}
			}
		j++
		}
	i++
	}

	Println(len(u) + len(w))

	for l < 2 {
		Strings(w)
		for _, s := range w {
			Print(s)
		}
		w = u
	l++
	}
}