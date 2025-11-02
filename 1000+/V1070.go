package main
import . "fmt"
func main() {
	var (
		v                [1e3]float64
		a, b, p, q, r, s float64
		n, i, j          int
		P                = Println
	)

	Scan(&n)
	for i < n {
		Scan(&a, &b, &p, &q, &r)
		e := p*a*a + q*a + r
		x := a
		z := p*b*b + q*b + r
		if e > z {
			e = z
			x = b
		}
		z = -q / p / 2
		if z >= a && z <= b {
			p = p*z*z + q*z + r
			if e > p {
				e = p
				x = z
			}
		}
		s += e
		v[i] = x
		i++
	}

	P(s)
	for j < n {
		P(v[j])
		j++
	}
}