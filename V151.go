package main
import . "fmt"
func main() {
	var (
		o             [101]int
		N, M, l, j, i int
		s             = "NO"
		p             = 1
	)
	Scan(&N, &M)

	g := make([][]int, N+1)
	for l < M {
		a := 0
		b := 0
		Scan(&a, &b)
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
		l++
	}

	for j < N {
		j++
		o[j] = -1
	}

	for i < N && p > 0 {
		i++
		if o[i] < 0 {
			q := []int{i}
			o[i] = 0
			for len(q) > 0 && p > 0 {
				j = q[0]
				q = q[1:]
				for _, v := range g[j] {
					if o[v] < 0 {
						o[v] = o[j] ^ 1
						q = append(q, v)
					} else if o[v] == o[j] {
						p = 0
						break
					}
				}
			}
		}
	}

	if p > 0 {
		s = "YES"
	}
	Print(s)
}