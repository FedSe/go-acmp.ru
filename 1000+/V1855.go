package main
import . "fmt"
func main() {
	var a, b, c int
	Scan(&a, &b, &c)

	s := F(a + c - b, c + b - a, a + b - c)

	Print(s)
}

func F(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= a && b <= c {
		return b
	}
	return c
}