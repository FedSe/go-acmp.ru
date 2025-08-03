package main
import . "fmt"
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
		k := float64(t) / float64(i)
		for _, g := range s {
			if i > g.L {
				k += float64(g.H)
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