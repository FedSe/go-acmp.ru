package main
import (
	. "fmt"
	. "sort"
)

func G(a int) string {
	s := ""
	w := s
	if a == 0 {
		s = "0"
	}
	if a < 0 {
		w = "-"
		a = -a
	}
	for a > 0 {
		s = string(a%10+48) + s
		a = a / 10
	}
	s = w + s
	return s
}

func F(x []int, a int, b int) string {
	w := G(x[a]) + ", ..., " + G(x[b])
	v := ""
	z := a
	for z <= b {
		v = v + G(x[z]) + ", "
		z++
	}
	v = v[:len(v)-2]
	if len(w) <= len(v) {
		return w
	}
	return v
}

func main() {
	var (
		x       []int
		n, t, i int
		P       = Print
		a       = map[int]int{}
	)

	Scan(&n)
	for 0 < n {
		Scan(&t)
		a[t] = 1
		n--
	}

	for s := range a {
		x = append(x, s)
	}

	Ints(x)
	n = len(x)

	for i < n-2 && n > 2 {
		if x[i]+1 == x[i+1] && x[i+1]+1 == x[i+2] {
			k := i + 1
			for k <= n-2 && x[k]+1 == x[k+1] {
				k++
			}
			P(F(x, i, k))
			i = k + 1
			if i < n {
				P(", ")
			}
		} else {
			P(x[i], ", ")
			i++
		}
	}
	n--
	if i < n {
		P(x[i], ", ")
		i++
	}
	if i == n {
		P(x[i])
	}
}