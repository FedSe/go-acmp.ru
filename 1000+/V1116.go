package main
import . "fmt"
func main() {
	var h, m, s, x, y, z int
	Scan(&h, &m, &s, &x, &y, &z)
	Print(x*3600 + y*60 + z - h*3600 - m*60 - s)
}