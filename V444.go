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
		s = string(a%10+'0') + s
		a = a / 10
	}
	s = w + s
	return s
}

func F(x []int, a int, b int) string {
	w := G(x[a]) + ", ..., " + G(x[b])
	v := ""
	for
	z := a
	z <= b
	{
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
	var n, t, i int
	Scan(&n)
	a := make(map[int]bool)
	for 0 < n {
		Scan(&t)
		a[t] = true
	n--
	}
	x := make([]int, 0)
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
			Print(F(x, i, k))
			i = k + 1
			if i < n {
				Print(", ")
			}
		} else {
			Print(x[i], ", ")
			i++
		}
	}
n--
	if i < n {
		Print(x[i], ", ")
		i++
	}
	if i == n {
		Print(x[i])
	}
}