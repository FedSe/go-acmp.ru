package main
import (
	. "fmt"
	. "math"
)
func main() {
	var (
		n, i, j, x, y, z int
		m                = -1.
		a                [100]float64
		P                = Println
	)

	Scan(&n)
	for j < n {
		Scan(&a[j])
		j++
	}

	for i < n {
		i++
		j = i
		for j < n {
			j++
			k := j
			for k < n {
				d := a[i-1]
				b := a[j-1]
				c := a[k]
				p := (d + b + c) / 2
				p = Sqrt(p * (p - d) * (p - b) * (p - c))
				if d >= b+c || b >= d+c || c >= d+b {
					p = -1
				}
				k++
				if p > m {
					m = p
					x = i
					y = j
					z = k
				}
			}
		}
	}

	P(m)
	if m > -1 {
		P(x, y, z)
	}
}