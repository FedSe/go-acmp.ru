package main
import . "fmt"
func main() {
	var c, s [33]int
	n := 0
	z := 2
	c[0] = 1
	s[0] = 1

	Scan(&n)
	for z <= n {
		i := s[z-2]
		c[z] = c[z-2] + 2*i
		s[z] = i + c[z]
	z++
	}

	Print(c[n])
}