package main
import (
	. "fmt"
	. "sort"
)
func main() {
	var n, k, x, a, i int
	Scan(&n, &k, &x)

	b := make([]int, n)
	for i < n {
		Scan(&a)
		b[i] = 1
		m := k
		l := 0
		for l < 2 {
			s := 0
			f := a
			for f > 0 {
				s += f % m
				f /= m
			}
			b[i] *= s
			m = x
		l++
		}
	i++
	}
	Ints(b)
	
	for _, c := range b {
		Println(c)
	}
}