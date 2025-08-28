package main
import . "fmt"
func main() {
	var (
		o, l, z, n, m, r, q, g int
		x, d, y, h             [10][100]int
	)

	Scan(&m, &n)
	for l < m {
		j := 0
		for j < n {
			Scan(&r)
			x[l][j] = r
			j++
		}
		l++
	}

	for g < m {
		l = 0
		for l < n {
			y[g][l] = -1
			l++
		}
		g++
	}

	for z < m {
		d[z][0] = x[z][0]
		z++
	}
	z = 1
	for z < n {
		g = 0
		for g < m {
			d[g][z] = 1 << 30
			l = -1
			v := 1 << 30
			for _, k := range []int{g - 1, g, g + 1} {
				if k >= 0 && k < m {
					c := d[k][z-1]
					if c < v || (c == v && k < l && l != -1) {
						v = c
						l = k
					}
				}
			}
			if l > -1 {
				d[g][z] = v + x[g][z]
				y[g][z] = l
			}
			g++
		}
		z++
	}

	z = d[0][n-1]
	l = 1
	for l < m {
		if d[l][n-1] < z {
			z = d[l][n-1]
			o = l
		}
		l++
	}

	l = n
	for l > 0 {
		l--
		h[0][l] = o + 1
		o = y[o][l]
	}

	for q < n {
		Println(h[0][q])
		q++
	}
	Print(z)
}