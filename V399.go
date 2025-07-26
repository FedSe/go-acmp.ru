package main
import (
	b "bufio"
	. "fmt"
	. "os"
)
const (
	w int = iota
	R
	U
	L
)
func main() {
	var (
		a [101][101]int
		s [101]string
		n, m, p, v, i int
		u = w
		r = b.NewScanner(Stdin)
		x = 2
		y = 2
		o = 1
	)

	Scan(&n, &m)

	for r.Scan() {
		s[i] = "@" + r.Text()
	i++
	}
	s[0] = s[0][1:]

	a[2][2] = 1

	for !(x == m-1 && y == n-1) {
		var d, f, k int
		z := 2000000

		if s[y+1][x] != 64 {
			if a[y+1][x] < z {
				z = a[y+1][x]
				k = w
				d = 0
				f = 1
			}
		}

		if s[y][x+1] != 64 {
			if a[y][x+1] < z {
				z = a[y][x+1]
				k = R
				d = 1
				f = 0
			}
		}

		if s[y-1][x] != 64 {
			if a[y-1][x] < z {
				z = a[y-1][x]
				k = U
				f = -1
				d = 0
			}
		}

		if s[y][x-1] != 64 {
			if a[y][x-1] < z {
				z = a[y][x-1]
				k = L
				d = -1
				f = 0
			}
		}

		if p > 9999999 || z == 2000000 {
			p = -1
			break
		}

		p++
		if k == u {
			x += d
			y += f
			a[y][x]++
		} else {
			if s[y+o][x+v] != 64 && a[y+f][x+d] == a[y+o][x+v] {
				x += v
				y += o
				k = u
				a[y][x]++
			} else {
				u = k
				x += d
				y += f
				a[y][x]++
				v = d
				o = f
			}
		}
	}

	Print(p)
}