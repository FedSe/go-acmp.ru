package main
import . "fmt"
func main() {
	var n, i, r, j int
	Scan(&n)
	s := ""

	for n > 0 {
		s = string(n % 2 + 48) + s
		n /= 2
	}

	m := s
	for i < len(s) {
		s = s[len(s)-1:] + s[0:len(s)-1]
		if s > m {
			m = s
		}
	i++
	}

	for j < len(m) {
		r *= 2
		r += int(m[j] - 48)
	j++
	}

	Print(r)
}