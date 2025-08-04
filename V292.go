package main
import . "fmt"
func G(n int) int {
	if n < 10 {
		switch n {
		case 2, 3, 5, 7:
			return n
		}
		return 0
	}
	a := 0
	if n%2 > 0 {
		a = n - 1
		if n > 2 {
			i := 3
			a = n
			for i*i <= n {
				if n%i < 1 {
					a = 0
				}
				i += 2
			}
		}
	}
	if a > 0 {
		return n
	}

	s := 0
	for n > 0 {
		s += n % 10
		n /= 10
	}

	return G(s)
}

func main() {
	n := 0
	Scan(&n)
	Print(G(n))
}