package main
import . "fmt"
func main() {
	var (
		f, t                [100]string
		n, m, a, o, i, j, x int
		S                   = Scan
	)

	S(&n, &m)
	for i < n {
		S(&f[i])
		i++
	}

	S(&o, &x)
	for j < o {
		S(&t[j])
		j++
	}

	x -= m
	o -= n
	for 0 <= x {
		y := 0
		for y <= o {
			i = x
			for i < x+m {
				j = y
				for j < y+n {
					if f[j-y][i-x]-t[j][i] > 200 {
						goto A
					}
					j++
				}
				i++
			}
			a++
		A:
			y++
		}
		x--
	}

	Print(a)
}