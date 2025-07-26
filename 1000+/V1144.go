package main
import . "fmt"
func main() {
	var k, n, m, x, y, z int
	Scan(&x, &y, &z)
	if x*y > 0 {
		for z > 0 {
			k++
			if y > z && x < y {
				n++
				if n == 2 || n > 2 && k < m {
					m = k
				}
			k = 0
			}
			x = y
			y = z
			Scan(&z)
		}
	}
	Print(m)
}