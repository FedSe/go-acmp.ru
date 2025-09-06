package main
import . "fmt"
func main() {
	var n, k, s, b int
	r := 1

	Scan(&n)
	for n > 0 {
		z := n % 10
		n /= 10
		a := k
		i := k
		for i > 1 {
			a *= 10
			i--
		}
		s += z*a + r
		if z < 6 {
			s -= r
			if z == 5 {
				s += b + 1
			}
		}
		b += z * r
		r *= 10
		k++
	}

	Print(s)
}