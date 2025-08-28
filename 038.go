package main
import . "fmt"
type T [100]int
var (
	w, y [100]T
	a    T
	e, n int
)
func H(t, h, p int) int {
	f := &w[t][h]
	g := &y[t][h]
	if *g < 1 {
		k := a[t]
		if p < 1 {
			k = -k
		}
		x := k
		if t == h {
			return x
		}
		x += H(t+1, h, -p)
		k = a[h]
		if p < 1 {
			k = -k
		}
		z := H(t, h-1, -p) + k
		*g = 1
		*f = z
		if p == 1 == (x > z) {
			*f = x
		}
	}
	return *f
}

func main() {
	Scan(&n)
	for e < n {
		Scan(&a[e])
		e++
	}

	e = H(0, n-1, 1)
	if e > 0 {
		e = 1
	}
	if e < 0 {
		e = 2
	}

	Print(e)
}