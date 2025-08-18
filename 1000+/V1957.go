package main
import . "fmt"
func P(x int) int {
	n := x
	s := 0
	a := 0

	for n > 0 {
		a = a*10 + n%10
		n /= 10
	}

	n = x
	for n > 0 {
		s += n % 10
		n /= 10
	}
	if s > 10 {
		k := x
		if x > 9 {
			k %= 100
		}
		a = x*100 + k
	}
	return a
}
func main() {
	m := 1
	n := 1

	Scan(&m)
	for {
		if P(P(n)) >= m {
			Print(n)
			break
		}
		n++
	}
}