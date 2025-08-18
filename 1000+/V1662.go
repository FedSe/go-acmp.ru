package main
import (
	. "fmt"
	. "math"
)
func main() {
	var x, y, R float64
	P := Println

	Scan(&x, &y, &R)

	R = R*R - y*y
	if R < 0 {
		P(0)
	}
	if R == 0 {
		P(1, x)
	}
	if R > 0 {
		R = Sqrt(R)
		P(2, x-R, x+R)
	}
}