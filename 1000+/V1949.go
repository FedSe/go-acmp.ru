package main
import . "fmt"
func main() {
	n := 0
	c := 0
	d := 1

	Scan(&n)
	n++
	for d*d <= n {
		if n%d < 1 {
			c++
			if d*d != n {
				c++
			}
		}
		d++
	}

	Print(c)
}