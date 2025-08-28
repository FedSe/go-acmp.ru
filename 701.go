package main
import . "fmt"
func main() {
	n := 0
	r := 0
	p := 1
	s := ""

	Scan(&n, &s)
	n = n%10 + 2
	i := len(s)
	for i > 0 {
		i--
		q := int(s[i] - 48)
		a := p * (q - 7)
		if s[i] > 47 && s[i] < 58 {
			a = p * q
		}
		r += a
		p *= n
	}

	Print(r)
}