package main
import . "fmt"
func main() {
	var (
		q                      [101][101][101]int
		n, a, b, c, d, e, f, g int
	)

	Scan(&n)
	for 0 < n {
		Scan(&b, &c, &d, &e, &f, &g)
		for b < e {
			y := c
			for y < f {
				z := d
				for z < g {
					q[b][y][z] = 1
					z++
				}
				y++
			}
			b++
		}
		n--
	}

	for n < 101 {
		b = 0
		for b < 101 {
			c = 0
			for c < 101 {
				a += q[n][b][c]
				c++
			}
			b++
		}
		n++
	}

	Print(a)
}