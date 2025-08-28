package main
import (
	. "fmt"
	. "sort"
)
func main() {
	var (
		d       []int
		N, K, n int
		i       = 9
		P       = Print
		a       = "NO"
	)
	Scan(&N, &K)
	if K < 2 {
		a = "YES"
		P(a)
		return
	}
	t := K
	for i > 1 {
		for t%i < 1 {
			d = append(d, i)
			t /= i
		}
		i--
	}
	if t == 1 {
		Ints(d)
		for _, v := range d {
			n = n*10 + v
		}
		if n <= N && n > 0 {
			a = "YES"
		}
	}
	P(a)
}