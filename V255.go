package main
import (
	. "fmt"
	. "math"
)
func main() {
	n := 0
	Scan(&n)

	d := 1
	q := 2
	for q < int(Trunc(Sqrt(float64(n)))) + 2 {
		if n%q < 1 {
			d = n / q
			break
		}
	q++
	}

	Print(d, n-d)
}