package main
import . "fmt"
func main() {
	var (
		d          [100][100]int
		n, w, l, i int
		F          = 1 << 40
		s          = "NO"
	)

	Scan(&n)
	for l < n {
		j := 0
		for j < n {
			Scan(&w)
			if w == 1e5 {
				w = F
			}
			d[l][j] = w
			j++
		}
		l++
	}

	for i < n {
		w = 0
		for w < n {
			l = 0
			for l < n {
				q := d[w][i]
				p := d[i][l]
				if q < F && p < F && q+p < d[w][l] {
					d[w][l] = q + p
				}
				l++
			}
			w++
		}
		i++
	}

	for 0 < n {
		n--
		if d[n][n] < 0 {
			s = "YES"
		}
	}

	Print(s)
}