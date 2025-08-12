package main
import . "fmt"
func main() {
	r := "NO SOLUTION"
	s := r
	i := 0

	Scan(&s)
	n := []rune(s)
	j := len(n)
	for i < j {
		j--
		n[i], n[j] = n[j], n[i]
		i++
	}
	i = 1
	for i < len(s) {
		if s[i] != s[0] {
			r = s
			if s == string(n) {
				r = s[1:]
			}
			break
		}
		i++
	}

	Print(r)
}