package main
import (
	. "fmt"
	. "sort"
)

func main() {
	var (
		d []int
		l, s, n int
	)

	Scan(&n)

	for n > 0 {
		d = append(d, n%10)
		n /= 10
	}
	Ints(d)

	for i := range d {
		l = l*10 + d[len(d)-i-1]
	}

	if d[0] < 1 {
	    n = 1
		for d[n] < 1 {
			n++
		}
		d[0], d[n] = d[n], d[0]
	}

	for _ , i := range d {
		s = s*10 + i
	}

	Print(s, l)
}