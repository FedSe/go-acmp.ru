package main
import . "fmt"
func main() {
	n := 0
	i := 1
	s := "1"
	c := s

	Scan(&n)
	for len(s) <= n {
		s += c
		i++
		c += Sprint(i)
	}

	Print(s[n] - 48)
}