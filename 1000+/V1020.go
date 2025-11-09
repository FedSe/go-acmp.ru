package main
import . "fmt"
func main() {
	var n, x, i int
	y := 1 << 62

	Scan(&n)
	for i < 63 {
		if 1<<i&n > 0 {
			p := 1 << i
			if p < y {
				y = p
			}
			if p > x {
				x = p
			}
		}
		i++
	}

	Print(y, x)
}