package main
import (
	. "fmt"
	. "sort"
)
func main() {
	var (
		t       [12]int
		v, z, i int
		s       = map[int]int{}
		a       = "IMPOSSIBLE"
	)

	for i < 12 {
		Scan(&t[i])
		s[t[i]] = 1
		i++
	}

	for z < 11 {
		if t[z] == t[z+1] {
			v++
		}
		z += 2
	}

	Ints(t[:])
	z = len(s)

	if z == 1 || z == 2 && v == 2 && (t[0] == t[7] || t[11] == t[4]) || (t[0] == t[3] && t[4] == t[7] && t[8] == t[11] && z == 3 && v < 1) {
		a = a[2:]
	}

	Print(a)
}