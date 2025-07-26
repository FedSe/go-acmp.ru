package main
import . "fmt"
func main() {
	var h, m, s, x, y, z int
	Scan(&h, &m, &s, &x, &y, &z)
	h = h*3600 + m*60 + s
	x = x*3600 + y*60 + z
	s = x-h
	if h > x {
		s = h-x
	}
	Print(s)
}