package main
import (
	. "fmt"
	. "math"
)
func main() {
	var a, b, c, d, i float64
	A := Abs

	Scan(&a, &b, &c, &d)
	x := A(c - a)
	y := A(d - b)
	if int(x+y)&1 < 1 && (c != a || d != b) {
		if x == y {
			i++
			goto B
		}
		for i < 8 {
			i++
			j := 0.
			for j < 8 {
				j++
				if A(i-a) == A(j-b) && A(c-i) == A(d-j) {
					Print(2, i, j)
					return
				}
			}
		}
	}
B:
	Print(i)
}