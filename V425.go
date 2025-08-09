package main
import . "fmt"
func main() {
	var N, w, e, q, i int
	Scan(&N, &w, &e)
	e -= w
	for i < N {
		i++
		j := 0
		for j < N {
			j++
			c := 100 * i
			d := 100 * j
			a := c - 100
			b := d - 100
			n := N * 100
			for _, v := range [][4]int{
				{a, b, c, b},
				{c, b, c, d},
				{c, d, a, d},
				{a, d, a, b}} {
				x := v[0]
				p := v[1]
				k := v[2] - x
				m := v[3] - p
				p -= w
				h := float64(n*m - e*k)
				t := float64(x*m-p*k) / h
				u := float64(x*e-p*n) / h
				if t >= 0 && t <= 1 && u >= 0 && u <= 1 {
					q++
					break
				}
			}

		}
	}

	Print(q)
}