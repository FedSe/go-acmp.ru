package main
import . "fmt"
func main() {
	var x, y, z, n, p int
	a := ""
	Scan(&n)

	for 0 < n {
		Scan(&a, &p)
		if a == "X" { x += p }
		if a == "Y" {
			x += p
			z += p
		}
		if a == "Z" { z += p }
	n--
	}

	if x > 0 && z > 0 {
		y = x
		if y > z { y = z }
	}

	if x < 0 && z < 0 {
		y = -min(-x, -z)
	}

	x -= y
	z -= y

	if x < 0 { x = -x }
	if y < 0 { y = -y }
	if z < 0 { z = -z }
	Print(x + y + z)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}