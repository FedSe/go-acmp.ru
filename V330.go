package main
import . "fmt"
func main() {
	var x, y, z, c, a int

	Scan(&x, &y, &z, &c)
	x-=z
	y-=c
	if x < 0 { x = -x }
	if y < 0 { y = -y }

	if (x + y) % 2 < 1 {
		a = 2
	}

	if x == y {
		a = 1
	}

	Print(a)
}