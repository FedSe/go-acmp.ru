package main
import . "fmt"
type S struct {
	u, y int
}
func G(a []int, o, p int) int {
	k := a[o]
	if p < 1 {
		k = -k
	}
	return k
}
func H(a []int, t, h, p int, w [][]S) int {
	if w[t][h].y < 1 {
		x := G(a, t, p)
		if t == h {
			return x
		}
		x += H(a, t+1, h, -p, w)
		z := H(a, t, h-1, -p, w) + G(a, h, p)
		w[t][h].y = 1
		w[t][h].u = z
		if (p == 1) == (x > z) {
			w[t][h].u = x
		}
	}
	return w[t][h].u
}

func main() {
	var n, e, i int
	Scan(&n)
	a := make([]int, n)
	w := make([][]S, n)
	for e < n {
		Scan(&a[e])
	e++
	}
	for i < n {
		w[i] = make([]S, n)
	i++
	}
	e = H(a, 0, n-1, 1, w)
	if e > 0 {
		e = 1
	}
	if e < 0 {
		e = 2
	}
	Print(e)
}