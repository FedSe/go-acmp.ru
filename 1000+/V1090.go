package main
import . "fmt"
func main() {
	x := 0
	y := 0
	P := Println

	Scan(&x, &y)
	if x >= 0 && y >= 0 {
		P(1)
	}
	if x < 1 && y >= 0 {
		P(2)
	}
	if x < 1 && y < 1 {
		P(3)
	}
	if x >= 0 && y < 1 {
		P(4)
	}
}