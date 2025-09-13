package main
import (
	. "fmt"
	. "math"
)
func main() {
	var x, y, a int

	Scan(&x)
	for a < int(Sqrt(float64(x))) {
		a++
		if x%a < 1 {
			w := 0
			I := 0
			b := x / a
			for I < 2 {
				s := 0
				for b > 0 {
					s += b % 10
					b /= 10
				}
				I++
				w += s
				b = a
			}
			if w > y {
				y = w
			}
		}
	}

	Print(y)
}