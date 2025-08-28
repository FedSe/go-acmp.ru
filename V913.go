package main
import . "fmt"
type T = float64
func main() {
	var (
		s             [1e3][2]int
		n, m, t, i, z int
		w             = 1e18
	)

	Scan(&n, &m)
	for i < n {
		Scan(&z, &s[i][0], &s[i][1])
		t += z
		i++
	}

	n = 1
	i = 1
	for i <= m {
		k := T(t) / T(i)
		for _, g := range s {
			if i > g[0] {
				k += T(g[1])
			}
		}
		if k < w || k == w && i > n {
			w = k
			n = i
		}
		i++
	}

	Print(n)
}