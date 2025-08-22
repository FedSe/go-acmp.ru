package main
import . "fmt"
func main() {
	var (
		q       [8][8]int
		d       [8]int
		F       func(x int)
		n, l, k int
		w       = -1
	)

	Scan(&n)
	for l < n {
		j := 0
		for j < n {
			Scan(&q[l][j])
			j++
		}
		l++
	}

	z := make([]int, n)
	g := make([]int, n)

	F = func(x int) {
		i := 0
		if x == n {
			u := 0
			for i < n {
				j := 0
				for j < n {
					if j != g[i] {
						u += q[i][j]
					}
					j++
				}
				i++
			}
			if w == -1 || u < w {
				w = u
				copy(z, g)
			}
			return
		}
		for i < n {
			if d[i] < 1 {
				d[i] = 1
				g[x] = i
				F(x + 1)
				d[i] = 0
			}
			i++
		}
	}

	F(0)
	for k < n {
		Printf("%c", 65+z[k])
		k++
	}

	Print(`
`, w)
}