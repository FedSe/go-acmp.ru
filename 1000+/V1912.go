package main
import . "fmt"
func main() {
	x := 0
	y := 0

	Scan(&x, &y)
	x -= y
	if x < 0 { x = -x }

	Print((x + 4) / 5)
}