package main
import (
	. "fmt"
	. "sort"
)
func main() {
	var n, z, a, t int
	Scan(&n)

	p := make([]int, n)
	for z < n {
		Scan(&p[z])
	z++
	}

	if n < 3 {
		for 0 < n {
		n--
			t += p[n]
		}
		Print(t)
		return
	}

	Ints(p)
	n--
	a = n - 2
	t = p[n] + p[n-1]
	m := t
	for a >= 0 {
		for a >= 0 && p[a]+p[a+1] < p[n] {
			if t > m {
				m = t
			}
			t -= p[n]
			n--
			t += p[a]
			a--
		}
		if a >= 0 {
			t += p[a]
			if t > m {
				m = t
			}
			a--
		}
	}

	Print(m)
}