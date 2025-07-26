package main
import ."fmt"
func main() {
	n:=0
	Scan(&n)
	f, g, h := 1, 1, 1

	for 0 < n {
		f = (h+g) % 10
		h = g
		g = f
	n--
	}

	Print(h)
}