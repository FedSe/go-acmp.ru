package main
import . "fmt"
func main() {
	var x, y, z, v, b, n float64
	Scan(&x, &y, &z, &v, &b, &n)

	x -= z
	n -= v
	y -= v
	b -= z
	z = x*n - y*b
	if z < 0 {
		z = -z
	}

	Print(z / 2)
}