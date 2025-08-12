package main
import . "fmt"
func main() {
	n := 0
	i := 0

	Scan(&n)
	d := n
	m := n
	for i < 4e4 {
		i++
		var (
			s [10]int
			a = n - i
			x = i
			t = 0
		)
		if a < 0 {
			a = -a
		}
		for x > 0 {
			if s[x%10] < 1 {
				t++
				s[x%10] = 1
			}
			x /= 10
		}
		if t < 3 && a < d {
			d = a
			m = i
		}
	}

	Print(m)
}