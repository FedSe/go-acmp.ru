package main
import . "fmt"
type S struct {
	u, y int
}
var (
	w [100][100]S
	a [100]int
)
func H(t, h, p int) int {
	f := &w[t][h]
	if f.y < 1 {
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
		f.y = 1
		f.u = z
		if p == 1 == (x > z) {
			f.u = x
		}
	}
	return f.u
}

func main() {
	n := 0
	e := 0

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