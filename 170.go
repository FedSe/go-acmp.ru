package main
import (
	. "fmt"
	. "math"
)
func main() {
	n := 0
	a := 1

	Scan(&n)
	k := int(Sqrt(2 * float64(n)))
	for k*k+k > 2*n {
		k--
	}
	for k > 0 {
		if 2*n%k < 1 {
			t := 2*n/k - k
			if t > 0 && t%2 > 0 {
				a = k
				break
			}
		}
		k--
	}

	Print(a)
}