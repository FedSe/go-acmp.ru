package main
import . "fmt"
type S struct{ t, v int }
func main() {
	var (
		s                      [101]int
		g                      [101][][3]int
		N, d, v, a, c, b, f, i int
	)

	Scan(&N, &d, &v, &i)
	for 0 < i {
		Scan(&a, &c, &b, &f)
		g[a] = append(g[a], [3]int{b, c, f})
		i--
	}

	for i < N {
		i++
		s[i] = 10001
	}
	s[d] = 0

	z := []S{{0, d}}
	for len(z) > 0 {
		o := z
		d = len(o) - 1
		u := o[d]
		z = o[:d]
		if u.t <= s[u.v] {
			for _, e := range g[u.v] {
				i = e[0]
				d = e[2]
				if u.t <= e[1] && d < s[i] {
					s[i] = d
					z = append(z, S{d, i})
				}
			}
		}
	}

	if s[v] > 10000 {
		s[v] = -1
	}
	Print(s[v])
}