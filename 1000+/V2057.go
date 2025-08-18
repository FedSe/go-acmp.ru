package main
import (
	. "fmt"
	. "strings"
)
func F(s string) [4]int {
	var r [4]int
	j := 0
	v := 0
	for j < 4 {
		Sscan(Split(s, ".")[j], &v)
		r[j] = v
		j++
	}
	return r
}

func main() {
	var (
		x, y, i int
		a       = ""
		m       = a
	)

	Scan(&a, &m)
	p := F(a)
	q := F(m)

	for i < 4 {
		x = (x << 8) | p[i]
		y = (y << 8) | q[i]
		i++
	}

	Print(x &^ y)
}