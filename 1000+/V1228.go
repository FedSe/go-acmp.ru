package main
import . "fmt"
func F(n int) int {
	a := 0
	if n > 1 {
		k := n / 2
		for n%k > 0 {
			k--
		}
		if k < 2 {
			a = n
		}
	}
	return a
}

func main() {
	var a, b, c int
	w := " No"

	Scan(&a, &b, &c)

	s := F(a) + F(b) + F(c)
	if F(s) > 0 {
		w = " Yes"
	}
	Print(s, w)
}