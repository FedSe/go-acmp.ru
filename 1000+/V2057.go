package main
import (
	. "fmt"
	. "strings"
)
func main() {
	var (
		p             [8]int
		x, y, i, v, j int
		a             = ""
		m             = a
	)

	Scan(&a, &m)
	a += "." + m
	for j < 8 {
		Sscan(Split(a, ".")[j], &v)
		p[j] = v
		j++
	}

	for i < 4 {
		x = x<<8 | p[i]
		y = y<<8 | p[i+4]
		i++
	}

	Print(x &^ y)
}