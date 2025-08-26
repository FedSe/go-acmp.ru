package main
import . "fmt"
type T = float64
func main() {
	var n, m, t, i, z int

	Scan(&n, &m)
	s := make([]struct{ L, H int }, n)

	for i < n {
		Scan(&z, &s[i].L, &s[i].H)
		t += z
		i++
	}

	w := 1e18
	n = 1
	i = 1
	for i <= m {
		k := T(t) / T(i)
		for _, g := range s {
			if i > g.L {
				k += T(g.H)
			}
		}
		if k < w || (k == w && i > n) {
			w = k
			n = i
		}
		i++
	}

	Print(n)
}