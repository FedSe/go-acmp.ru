package main
import . "fmt"
func main() {
	a := "NO"
	s := a
	t := s
	j := 0
	i := 0

	Scan(&s, &t)
	for i < len(t) {
		if s[j] == t[i] {
			j++
		}

		if j == len(s) {
			a = "YES"
			t = ""
		}
		i++
	}

	Print(a)
}