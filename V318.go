package main
import (
	. "fmt"
	. "math"
)
func main() {
	var (
		r [32]int
		n, i, k, s, z, j int
	)
	Scan(&n)

	if n > 0 {
		for n > 0 {
			r[i] = n % 2
			n /= 2
			i++
		}

		for r[k] < 1 {
			k++
		}
		r[k] = 0
		k++

		for r[k] > 0 {
			r[k] = 0
			k++
			s++
		}
		r[k] = 1
		for j < s {
			r[j] = 1
		j++
		}

		for z < 32 {
			n += r[z] * int(Pow(2, float64(z)))
		z++
		}
	}

	Print(n)
}