package main
import (
	. "fmt"
	. "sort"
)
func main() {
	var n, k, x, a int
	Scan(&n, &k, &x)

	b := make([]int, n)
	for 0 < n {
		n--
		Scan(&a)
		b[n] = 1
		m := k
		l := 0
		for l < 2 {
			s := 0
			f := a
			for f > 0 {
				s += f % m
				f /= m
			}
			b[n] *= s
			m = x
			l++
		}
	}
	Ints(b)

	for _, c := range b {
		Println(c)
	}
}