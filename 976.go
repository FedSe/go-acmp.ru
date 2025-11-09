package main
import . "fmt"
var n, a int
func F(i, m, p, s int) {
	if i == n {
		if p == s {
			a++
		}
	} else if p <= s+n-i {
		for m > 0 {
			F(i+1, m, p*m, s+m)
			m--
		}
	}
}

func main() {
	Scan(&n)
	F(0, n, 1, 0)
	Print(a)
}