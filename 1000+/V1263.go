package main
import . "fmt"
func main() {
	var n, q, x int
	s := ""

	Scan(&n)
	for n > 0 {
		Scan(&s)
		z := 0
		switch s[0] {
		case 66, 77, 83:
			z = 1
		case 68, 71, 74, 75, 84, 87:
			z = 2
		}
		q -= z
		if q < 0 {
			q = -q
		}
		x += q
		q = z
		n--
	}

	Print(x)
}