package main
import (
	. "fmt"
	. "sort"
	. "os"	
	b "bufio"
)
type M struct {
	t, w, d, s int
}

func main() {
	var i, k, n int

	r := b.NewScanner(Stdin)
	Scanln(&n)

	a := make([]M, n)

	for r.Scan() {
		Sscan(r.Text(), &a[i].t, &a[i].w)
		a[i].d = i
	i++
	}

	Slice(a, func(i, j int) bool {
		return a[i].w < a[j].w
	})

	a[0].s = a[0].t * a[0].w

	for
	i := 1
	i < n
	{
		c := a[i].t * a[i].w
		a[i].s = a[i-1].s
		if a[i].s < c {
			a[i].s = c
		}
	i++
	}

	Slice(a, func(i, j int) bool {
		return a[i].d < a[j].d
	})

	for k < n {
		Println(a[k].s)
	k++
	}
}