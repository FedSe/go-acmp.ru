package main
import (
	. "fmt"
	. "sort"
)
func main() {
	var (
		w, s, u []int
		a       [1e5]int
		q       = map[int]int{}
		m       = map[int]int{}
		N, i    int
		P       = Print
	)

	Scan(&N)
	for i < N {
		Scan(&a[i])
		q[a[i]] = 1
		i++
	}

	for x := range q {
		u = append(u, x)
	}
	Ints(u)

	i = len(u)
	if i < 2 {
		P("Yes 1")
		return
	}

	if i == 3 {
		i = u[1]
		s = append(s, i-u[0])
		if i-u[0] != u[2]-i {
			i = 0
		}
	}
	if i > 0 {
		if i == 2 {
			i = u[1] - u[0]
			s = append(s, i)
			if i&1 < 1 {
				s = append(s, i/2)
			}
		}

		for _, x := range s {
			if x > m[x] {
				m[x] = 1
				w = append(w, x)
			}
		}
		Ints(w)

		for _, X := range w {
			i = 0
			for _, T := range []int{-X, 0, X} {
				v := 1
				for _, y := range a[:N] {
					d := y - a[0] - T
					if d < 0 {
						d = -d
					}
					if d > 0 && d != X {
						v = 0
						break
					}
				}
				if v > 0 {
					i = 1
					break
				}
			}
			if i > 0 {
				P("Yes ", X)
				return
			}
		}
	}

	P("No")
}