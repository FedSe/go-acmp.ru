package main
import . "fmt"
var k, n int

func F(n int) int {
	a := 0
	if n > 0 {
		a = F(n-1) + k
		a %= n
	}
	return a
}

func main() {
	Scan(&n, &k)

	Print(F(n)+1)
}