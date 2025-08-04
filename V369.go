package main
import . "fmt"
type P struct{ a, b int }
func main() {
	var (
		Q, W, E, e                [101]int
		N, K, T, a, l, j, i, k, z int
		S                         = Scan
	)

	S(&N, &K, &T)
	for l < N {
		S(&Q[l])
		l++
	}
	for j < N {
		S(&W[j])
		j++
	}
	for z < N {
		S(&E[z])
		z++
	}

	w := make([][]P, T+1)
	for i < N {
		if Q[i] <= T {
			w[Q[i]] = append(w[Q[i]], P{E[i], W[i]})
		}
		i++
	}

	for k := range e {
		e[k] = -1
	}
	e[0] = 0
	z = 1
	for z <= T {
		var u [101]int
		for k := range u {
			u[k] = -1
		}
		i = 0
		for i <= K {
			l = -1
			if i-1 >= 0 && e[i-1] > -1 {
				j = e[i-1]
				if l < j {
					l = j
				}
			}
			if e[i] > -1 {
				j = e[i]
				if l < j {
					l = j
				}
			}
			if i+1 <= K && e[i+1] > -1 {
				j = e[i+1]
				if l < j {
					l = j
				}
			}
			if l > -1 {
				u[i] = l
			}
			i++
		}
		for _, g := range w[z] {
			l = g.a
			i = g.b
			if l >= 0 && l <= K {
				if u[l] > -1 {
					u[l] += i
				}
			}
		}
		e = u
		z++
	}

	for k <= K {
		if e[k] > a {
			a = e[k]
		}
		k++
	}
	Print(a)
}