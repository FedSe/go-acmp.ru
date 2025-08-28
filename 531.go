package main
import . "fmt"
func main() {
	var a, b, c, d, e, f, r, i int
	Scan(&a, &b, &c, &d, &e, &f, &r)

	l := f - r
	o := f + r
	F := func(x int) int {
		x += e
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
	w := F(0)

	y := r
	for i < r-1 {
		i++
		for i*i+y*y > r*r {
			y--
		}
		l = f - y
		o = f + y
		w += F(i) + F(-i)
	}

	l = f
	o = f
	w += F(r) + F(-r)

	Print(w)
}