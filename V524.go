package main
import (
	. "fmt"
	. "math"
)
func main() {
	var a, b, c, d, i float64

	Scan(&a, &b, &c, &d)
	x := Abs(c - a)
	y := Abs(d - b)
	if int(x+y)&1 < 1 && (c != a || d != b) {
		if x == y {
			i++
			goto A
		}
		for i < 8 {
			i++
			j := 0.
			for j < 8 {
				j++
				if Abs(i-a) == Abs(j-b) && Abs(c-i) == Abs(d-j) {
					Print(2, i, j)
					return
				}
			}
		}
	}
A:
	Print(i)
}