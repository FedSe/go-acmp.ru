package main
import . "fmt"
func main() {
	var x, p, y, k int
	Scan(&x, &p, &y)
	x *= 100
	y *= 100
	for x < y {
		k++
		x = int(float64(x) * (1 + float64(p)/100))
	}
	Print(k)
}