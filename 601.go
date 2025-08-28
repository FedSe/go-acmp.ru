package main
import . "fmt"
type T []int
func main() {
	var n, m, u, v, c int
	S := Scan

	S(&n, &m)
	g := make([][]T, n+1)
	for m > 0 {
		S(&u, &v, &c)
		g[u] = append(g[u], T{c, v})
		g[v] = append(g[v], T{c, u})
		m--
	}

	S(&m)
	v = 1
	for m > 0 {
		S(&c)
		u = -1
		for _, e := range g[v] {
			if e[0] == c {
				u = e[1]
				break
			}
		}
		if u < 0 {
			for m > 0 {
				S(&c)
				m--
			}
			Print("INCORRECT")
			return
		}
		v = u
		m--
	}

	Print(v)
}