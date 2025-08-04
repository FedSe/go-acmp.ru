package main
import (
	. "fmt"
	. "math/big"
)
func main() {
	var (
		d          [101][909]Int
		n, l, j, i int
		r          = &Int{}
	)

	Scan(&n)
	if n%2 < 1 {
		n /= 2
		m := n * 9
		d[0][0].SetInt64(1)
		for i < n {
			i++
			l = 0
			for l <= m {
				j := 0
				for j <= 9 {
					if l >= j {
						d[i][l].Add(&d[i][l], &d[i-1][l-j])
					}
					j++
				}
				l++
			}
		}

		for j <= m {
			w := d[n][j]
			r.Add(r, new(Int).Mul(&w, &w))
			j++
		}
	}

	Print(r)
}