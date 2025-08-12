package main
import . "fmt"
func main() {
	var (
		n, q, l, k, s, i, j int
		w                   [20][4]int
		S                   = Scan
	)

	S(&n)
	for i < n {
		z := 0
		for z < 4 {
			S(&w[i][z])
			z++
		}
		i++
	}

	S(&q, &l, &k, &i)
	for j < n {
		a := w[j][0]
		b := w[j][1]
		c := w[j][2]
		d := w[j][3]
		if a > c {
			a, c = c, a
		}
		if b > d {
			b, d = d, b
		}
		if q > k {
			q, k = k, q
		}
		if l > i {
			l, i = i, l
		}

		z := a
		if z < q {
			z = q
		}
		x := b
		if x < l {
			x = l
		}
		n := c
		if n > k {
			n = k
		}
		m := d
		if m > i {
			m = i
		}

		if n > z && m > x {
			s += (n - z) * (m - x)
		}
		j++
	}

	Print(s)
}