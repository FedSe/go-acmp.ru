package main
import . "fmt"
func main() {
	x := 0
	y := 0
	f := int(1e9)
	Scan(&x, &y)

	if x == f || x == -f || y == f || y == -f {
		Print("NO")
		return
	}

	z := x + 1
	if x < 0 {
		z -= 2
	}

	b := y + 1
	if y < 0 {
		b -= 2
	}

	Print(`YES
`, z, y, x, b, x*3-z-x, y*3-y-b)
}