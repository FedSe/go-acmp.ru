package main
import . "fmt"
func main() {
	var x, p, y, a int
	Scan(&x, &p, &y)

	x *= 100
	for x < y*100 {
		x += x * p / 100
		a++
	}

	Print(a)
}