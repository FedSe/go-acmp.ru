package main
import (
	. "fmt"
	. "sort"
)
func main() {
	var (
		e    []byte
		z    = make([][]int, 26)
		w    = make([][]int, 26)
		X    = ""
		Y    = X
		i, j int
	)
	Scan(&X, &Y)

	for i, c := range X {
		z[c-'a'] = append(z[c-'a'], i)
	}
	for i, c := range Y {
		w[c-'a'] = append(w[c-'a'], i)
	}

	for {
		u := 0
		c := 25
		for c >= 0 {
			a := F(z[c], i)
			b := F(w[c], j)
			if a != -1 && b != -1 {
				e = append(e, byte('a'+c))
				i = a + 1
				j = b + 1
				u = 1
				break
			}
			c--
		}
		if u < 1 {
			break
		}
	}

	if len(e) > 0 {
		Print(string(e))
	}
}

func F(p []int, s int) int {
	i := SearchInts(p, s)
	if i < len(p) {
		return p[i]
	}
	return -1
}