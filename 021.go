package main
import . "fmt"
func main() {
	var a, g, i, l int
	l = 1e5

	for i < 3 {
		Scan(&a)
		if a > g {
			g = a
		}
		if a < l {
			l = a
		}
		i++
	}

	Print(g - l)
}