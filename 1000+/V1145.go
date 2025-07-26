package main
import (
	. "fmt"
	. "math"
)
func main() {
	var a, b, c, x, n float64
	Scan(&x)
	for x > 0 {
		n++
		b += x
		c += x * x
		Scan(&x)
	}
	a = b / n
	Print(Sqrt((c + a*a*n - 2*b*a) / (n-1)))
}