package main
import . "fmt"

type A struct{ a, b, c int }

var (
	C, w, p, r, c, v             [42]int
	g                            [42][]A
	N, M, x, y, j, l, i, k, o, z int
	S                            = Scan
)

func F(u, t, f int) int {
	if u == t {
		return f
	}
	i := 0
	for i < len(g[u]) {
		e := &g[u][i]
		if e.b > 0 && v[e.a] >= v[u] {
			if e.b < f {
				f = e.b
			}
			k = F(e.a, t, f)
			if k > 0 {
				e.b -= k
				g[e.a][e.c].b += k
				return k
			}
		}
		i++
	}
	return 0
}

func E(h, t, p int) {
	g[h] = append(g[h], A{t, p, len(g[t])})
	if p > 0 {
		E(t, h, 0)
	}
}

func main() {
	S(&N, &M)

	for j < N {
		S(&r[j])
		j++
	}
	for l < M {
		S(&C[l])
		l++
	}

	for i < N {
		j = 0
		for j < M {
			S(&z)
			if z > 0 {
				l = z
				w[i] += l
				p[j] += l
				x += l
			}
			if p[j] > C[j] {
				x = -1
				goto A
			}
			c[j] = C[j] - p[j]
			j++
			if z < 0 {
				E(i+1, N+j, 1e9)
			}
		}
		r[i] -= w[i]
		i++
	}

	z = N + M + 1

	for k < N {
		E(y, k+1, r[k])
		k++
	}

	for o < M {
		E(N+1+o, z, c[o])
		o++
	}

	for v[z] > -1 {
		for i := range v {
			v[i] = -1
		}
		q := []int{y}
		v[y] = 0
		for len(q) > 0 {
			o = q[0]
			q = q[1:]
			for _, e := range g[o] {
				l = e.a
				if e.b > 0 && v[l] < 0 {
					v[l] = v[o] + 1
					q = append(q, l)
				}
			}
		}
		for {
			o = F(y, z, 1e9)
			if o < 1 {
				break
			}
			x += o
		}
	}
A:
	Print(x)
}