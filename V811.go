package main
import . "fmt"
func main() {
	var (
		b, d, e [101]string
		n, m, a, c, i, x, l, j int
	)

	Scan(&n, &m)
	for x < n {
		Scan(&b[x])
	x++
	}
	for i < n {
		x = 0
		for x < m/2+1 {
			d[i]   += "BW"
			d[i+1] += "WB"
			e[i]   += "WB"
			e[i+1] += "BW"
		x++
		}
	i += 2
	}
	for l < n {
		x = 0
		for x < m {
			s := "0"
			w := b[l][x]
			if d[l][x] == w {
				s = "1"
				a--
			}
			d[l] = d[l][:x] + s + d[l][x+1:]
			a++
			s = "0"
			if e[l][x] == w {
				s = "1"
				c--
			}
			e[l] = e[l][:x] + s + e[l][x+1:]
			c++
		x++
		}
	l++
	}

	if a < c {
		c = a
		e = d
	}

	Println(c)
	for j < n {
		x = 0
		for x < m {
			if e[j][x] < 49 {
				Println(j+1, x+1)
			}
		x++
		}
	j++
	}
}