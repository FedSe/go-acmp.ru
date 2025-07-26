package main
import . "fmt"
func main() {
	n := 0
	z := 0
	Scan(&n)

	q := make([]int, n)
	s := make([]byte, n)

	for z < n {
		q[z] = z
	z++
	}

	c := byte(66)

	for len(q) > 0 {
		s[q[0]] = c
		q = q[1:]
		c = 148 - c
		if len(q) > 0 {
			q = append(q, q[0])
			q = q[1:]
		}
	}

	Print(string(s))
}