package main
import . "fmt"
func F(a int) int {
	if a < 2 { return a }
	if a & 1 < 1 {
		return F(a/2)
	}
	return F(a/2) + F(a/2 + 1)
}

func main() {
	n := 0
	Scan(&n)

	Print(F(n))
}