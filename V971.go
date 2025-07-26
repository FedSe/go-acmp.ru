package main
import . "fmt"
func main() {
	var x, y, z, b int
	f := 1000000000
	Scan(&x, &y)

	if x == f || x == -f || y == f || y == -f {
		Print("NO")
		return
	}

	z = x + 1
	if x < 0 {
		z -= 2
	}

	b = y + 1
	if y < 0 {
		b -= 2
	}

	Print("YES\n", z, y,"\n", x, b,"\n", x*3 - z - x, y*3 - y - b)
}