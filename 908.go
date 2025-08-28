package main
import . "fmt"
var t = [1e6]int{0, 0, 1, 1, 2, 3, 2, 3, 3, 2, 3, 4, 3, 4, 4, 4, 4, 5, 3, 4, 4, 4, 5, 6, 4, 5, 5, 3, 4, 5, 4}
func F(n int) int {
	if t[n] > 0 {
		return t[n]
	}
	a := F(n / 3)
	if n%3 > 0 {
		a = F(n - 1)
		b := F(n / 2)
		if n&1 > 0 {
			b = F(n-2) + 1
		}
		if a > b {
			a = b
		}
	}
	return a + 1
}

func main() {
	n := 0
	z := 31
	for z < 1e6 {
		t[z] = F(z)
		z++
	}

	Scan(&n)

	Print(t[n])
}