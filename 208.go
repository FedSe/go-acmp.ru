package main
import . "fmt"
func main() {
	var n, i, r, j int
	s := ""

	Scan(&n)
	for n > 0 {
		s = Sprint(n%2) + s
		n /= 2
	}

	m := s
	for i < len(s) {
		n = len(s) - 1
		s = s[n:] + s[:n]
		if s > m {
			m = s
		}
		i++
	}

	for j < len(m) {
		r = r*2 + int(m[j]-48)
		j++
	}

	Print(r)
}