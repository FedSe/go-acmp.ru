package main
import (
	. "fmt"
	. "math"
	. "sort"
)

type A struct{ x, y float64 }

func F(o, a, b A) float64 {
	return (a.x-o.x)*(b.y-o.y) - (a.y-o.y)*(b.x-o.x)
}

func main() {
	var (
		o, w       []A
		p          [1e3]A
		n, i, l, k int
		e          = 0.
	)

	Scan(&n)
	for l < n {
		Scan(&p[l].x, &p[l].y)
		l++
	}

	Slice(p[:n], func(i, j int) bool {
		a := p[i]
		b := p[j]
		if a.x == b.x {
			return a.y < b.y
		}
		return a.x < b.x
	})

	for i < n {
		l = len(o) - 1
		for l > 0 && F(o[l-1], o[l], p[i]) < 1 {
			o = o[:l]
			l = len(o) - 1
		}
		o = append(o, p[i])
		i++
	}

	n--
	for n >= 0 {
		i = len(w) - 1
		for i > 0 && F(w[i-1], w[i], p[n]) < 1 {
			w = w[:i]
			i = len(w) - 1
		}
		w = append(w, p[n])
		n--
	}

	w = append(o[:l+1], w[:i+1]...)
	l = len(w)
	for k < l {
		b := w[(k+1)%l]
		e += Hypot(w[k].x-b.x, w[k].y-b.y)
		k++
	}

	Print(e)
}