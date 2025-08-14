package main
import (
	. "fmt"
	. "math"
)
func main() {
	n := 0
	d := 1
	q := 2

	Scan(&n)
	for q < int(Sqrt(float64(n)))+2 {
		if n%q < 1 {
			d = n / q
			break
		}
		q++
	}

	Print(d, n-d)
}