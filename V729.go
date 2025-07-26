package main
import . "fmt"
func main() {
	var (
		u, v [10]int
		n, m, a, b, s, j, i int
	)

	Scan(&n, &m)
	for s < n {
		Scan(&v[s])
	s++
	}
	if n < 2 {
		Print(1, v[0])
		return
	}
	for j < n {
		u[j] = j + 1
	j++
	}
	for 0 < m {
		Scan(&a, &b)
		s = u[a-1]
		h := u[b-1]
		if s > h {
			s, h = h, s
		}
		j = 0
		for j < n {
			if u[j] == s {
				u[j] = h
			}
		j++
		}
	m--
	}

	for i < n {
	i++
		j = 0
		s = 0
		for j < n {
			if u[j] == i {
				s += v[j]
			}
		j++
		}
		if s > 0 {
			Println(i, s)
		}
	}
}