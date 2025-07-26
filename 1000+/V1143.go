package main
import . "fmt"
func main() {
	var k, x, y int
	n := 1
	Scan(&y, &x)
	if x*y > 0 {
		for n > 0 {
			Scan(&n)
			if n > 0 && x > n && x > y {
				k++
			}
		y = x
		x = n
		}
	}
	Print(k)
}