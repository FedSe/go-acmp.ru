package main
import . "fmt"

func p(a, b, c, d, e, f, g, h int) int {
	s := 0
	if a > c {
		a, c = c, a
	}
	if b > d {
		b, d = d, b
	}
	if e > g {
		e, g = g, e
	}
	if f > h {
		f, h = h, f
	}

	z := a
	if z < e {
		z = e
	}
	x := b
	if x < f {
		x = f
	}
	n := c
	if n > g {
		n = g
	}
	m := d
	if m > h {
		m = h
	}

	if n > z && m > x {
		s = (n - z) * (m - x)
	}

	return s
}

func main() {
	var (
		n, b, c, d, s, i, j int
		a                   [20][4]int
		S = Scan
	)

	S(&n)
	for i < n {
		z := 0
		for z < 4 {
			S(&a[i][z])
			z++
		}
		i++
	}

	S(&b, &c, &d, &i)
	for j < n {
		s += p(a[j][0], a[j][1], a[j][2], a[j][3], b, c, d, i)
		j++
	}

	Print(s)
}