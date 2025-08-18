package main
import . "fmt"
func main() {
	var (
		g, d, p    [250][250]int
		n, i, j, l int
		s          = ""
		P          = Print
	)

	Scan(&n)
	for i < n {
		Scan(&s)
		j = 0
		for j < n {
			g[i][j] = int(s[j])
			j++
		}
		if i > 0 {
			d[0][i] = d[0][i-1] + g[0][i]
			d[i][0] = d[i-1][0] + g[i][0]
		}
		i++
	}

	i = 1
	for i < n {
		j = 1
		for j < n {
			q := d[i-1][j]
			h := d[i][j-1]
			if q < h {
				h = q
			}
			d[i][j] = h + g[i][j]
			j++
		}
		i++
	}

	i = n - 1
	j = n - 1
	for {
		p[i][j] = 1
		if i+j < 1 {
			break
		}
		if i < 1 {
			j--
		} else if j < 1 {
			i--
		} else {
			j--
			if d[i-1][j+1] < d[i][j] {
				i--
				j++
			}
		}
	}

	for l < n {
		j = 0
		for j < n {
			s = "."
			if p[l][j] > 0 {
				s = "#"
			}
			P(s)
			j++
		}
		P("\r\n")
		l++
	}
}