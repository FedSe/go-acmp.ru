package main
import . "fmt"
func main() {
	var c, h, o int

	Scan(&c, &h, &o)
	c /= 2
	h /= 6
	if c < o {
		o = c
	}
	if h < o {
		o = h
	}

	Print(o)
}