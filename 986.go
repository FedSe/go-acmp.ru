package main
import (
	. "fmt"
	. "math"
)
func main() {
	var n, q, w, i, x, y, l float64

	Scan(&n, &q, &w, &l)
	for n > 0 {
		Scan(&x, &y)
		i++
		y = w - y
		x = q - x
		if Sqrt(x*x+y*y) <= l {
			Print(i)
			return
		}
		n--
	}

	Print("Yes")
}