package main
import . "fmt"
func F(n, i int) int {
	r := 1
	t := n - i
	i = t
	for i < n {
		i++
		r *= i
		r /= i - t
	}
	return r
}

func main() {
	a := 0
	b := 0

	Scanf("%d:%d", &a, &b)
	if a < b {
		a, b = b, a
	}

	b = F(a+b-1, b)
	if a != 25 {
		b = F(48, 24) * 1 << (a - 26)
	}

	Print(b)
}