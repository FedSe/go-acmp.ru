package main
import (
	b "bufio"
	. "fmt"
	. "os"
	. "strconv"
)
var t [400000]int
func F(v, f, w int, a []int) {
	if f == w {
		t[v] = a[f]
		return
	}
	d := f + (w-f)/2
	F(2*v + 1, f, d, a)
	F(2*v + 2, d+1, w, a)
	t[v] = t[2*v + 1] + t[2*v + 2]
}
func H(v, f, w, l, r int) int {
	if r < f || l > w {
		return 0
	}
	if l <= f && w <= r {
		return t[v]
	}
	d := f + (w-f)/2
	return H(2*v + 1, f, d, l, r) + H(2*v + 2, d+1, w, l, r)
}
func main() {
	var n, m, l, r, i, j int
	s := b.NewScanner(Stdin)
	s.Split(b.ScanWords)
	Scan(&n)
	a := make([]int, n)
	for j < n {
		Scan(&a[j])
	j++
	}
	F(0, 0, n-1, a)
	Scan(&m)
	for s.Scan() {
		l, _ = Atoi(s.Text())
		s.Scan()
		r, _ = Atoi(s.Text())
		Println(H(0, 0, n-1, l-1, r-1))
	i++
	}
}