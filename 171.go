package main
import . "fmt"
func main() {
	x := 0
	Scan(&x)
	c := 1
	d := 2
	for d < 1e3 {
		e := 0
		for x%d < 1 {
			x /= d
			e++
		}
		if e > 0 {
			c *= e + 1
		}
		d++
	}
	Print(c)
}