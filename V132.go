package main
import . "fmt"
type T [101]int
func main() {
	var (
		s, d             T
		h                [101]T
		n, S, F, x, l, i int
		q                = 1 << 30
	)

	Scan(&n, &S, &F)
	for l < n {
		l++
		j := 0
		for j < n {
			j++
			Scan(&x)
			h[l][j] = x
		}
	}

	for i < n {
		i++
		d[i] = q
	}
	d[S] = 0

	for {
		l = -1
		i = 0
		for i < n {
			i++
			if s[i] < 1 && (l < 0 || d[i] < d[l]) {
				l = i
			}
		}
		if l < 0 || d[l] == q {
			break
		}
		s[l] = 1
		S = 0
		for S < n {
			S++
			i = h[l][S]
			if i > -1 {
				if d[l]+i < d[S] {
					d[S] = d[l] + i
				}
			}
		}
	}

	l = d[F]
	if l == q {
		l = -1
	}

	Print(l)
}