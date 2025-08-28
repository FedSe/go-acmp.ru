package main
import (
	. "fmt"
	. "sort"
	. "strings"
)
func main() {
	var (
		d          [255]int
		w          [255]string
		M, j, i, a int
	)

	Scan(&M)
	for j < M {
		Scan(&w[j])
		d[j] = 1
		j++
	}

	Slice(w[:M], func(i, j int) bool {
		return w[i] < w[j]
	})

	for i < M {
		j = 0
		for j < i {
			if len(w[j]) < len(w[i]) && HasPrefix(w[i], w[j]) && d[j]+1 > d[i] {
				d[i] = d[j] + 1
			}
			j++
		}
		if d[i] > a {
			a = d[i]
		}
		i++
	}

	Print(a)
}