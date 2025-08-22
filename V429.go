package main
import . "fmt"
func main() {
	var (
		v, l    [2e3]int
		c       [2e3][]int
		N, p, z int
		i       = 1
		t       = ""
		f       func(int, int) int
		S       = Scan
		P       = Print
	)

	S(&N)
	for i < N {
		i++
		S(&t, &p)
		c[p] = append(c[p], i)
		if t == "L" {
			S(&z)
			c[p] = append(c[p], i)
			v[i] = z
			l[i] = 1
		}
	}

	f = func(n, d int) int {
		if l[n] > 0 {
			return v[n]
		}
		r := -2
		if d&1 > 0 {
			r = 2
		}
		for _, v := range c[n] {
			x := f(v, d+1)
			if x > r == (d&1 < 1) {
				r = x
			}
		}
		return r
	}

	a := f(1, 0)
	if a > 0 {
		P("+")
	}

	P(a)
}