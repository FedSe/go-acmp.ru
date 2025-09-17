package main
import . "fmt"
type I = int

func F(a []I) map[I]I {
	r := map[I]I{}
	q := 1
	for range a {
		q *= 3
	}
	i := 0
	for i < q {
		w := 0
		z := 0
		t := i
		for _, v := range a {
			c := t % 3
			t /= 3
			w += c * v
			z += c
		}
		p, o := r[w]
		if !o || z < p {
			r[w] = z
		}
		i++
	}
	return r
}

func main() {
	var (
		A          [15]I
		n, m, s, i I
		P          = Print
	)

	Scan(&n, &m)
	for i < m {
		Scan(&A[i])
		s += 2 * A[i]
		i++
	}
	i = -1
	if s >= n {
		h := F(A[m/2 : m])
		i = 9e9
		for l, v := range F(A[:m/2]) {
			r, o := h[n-l]
			if o && v+r < i {
				i = v + r
			}
		}
		if i == 9e9 {
			i = 0
		}
	}

	P(i)
}