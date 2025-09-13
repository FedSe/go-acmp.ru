package main
import (
	. "fmt"
	. "math"
)
func main() {
	k := 0.
	l := 0.

	Scan(&k, &l)
	l -= 2 * k
	q := Sqrt(l*l - 4*k)
	z := (l + q) / 2
	l = (l - q) / 2
	a := l + 1
	b := z + 1
	if z*l == k {
		a, b = b, a
	}

	Print(int(a), b)
}