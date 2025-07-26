package main
import . "fmt"
func P(n int) {
	l := n % 3
	if l < 1 {
		l = 3
	}

	n -= l
	n /= 3

	if n > 0 {
		P(n)
	}

	Print(l)
}

func main() {
	n:=0
	Scan(&n)
	P(n)
}