package main
import . "fmt"

var a, b, c, d, e, f, r, i, l, o int

func F(x int) int {
	k := 0
	if !(x < a || x > c || l > d || o < b) {
		if l < b {
			l = b
		}
		if o > d {
			o = d
		}
		k = o - l + 1
	}
	return k
}

func main() {
	Scan(&a, &b, &c, &d, &e, &f, &r)
	l = f-r
	o = f+r
	w := F(e)

	y := r
	for i < r-1 {
	i++
		for i*i+y*y > r*r {
			y--
		}
		l = f-y
		o = f+y
		w += F(e+i) + F(e-i)
	}

	l = f
	o = f
	w += F(e+r) + F(e-r)

	Print(w)
}