package main
import (
	. "fmt"
	. "math"
)
func main() {
	n := 0.
	z := 26.
	Scan(&n)
	for z > 0 && n > 1 {
	z--
		a := Pow(2, z)
		if n > a {
			n -= a - 1
		}
		n--
	}
	Printf("%c", 96+int(z))
}