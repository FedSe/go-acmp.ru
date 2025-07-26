package main
import . "fmt"
func main() {
	n := 0
	r := 0
	p := 1
	s := ""
	Scan(&n, &s)
	n = n%10+2

	for
	i := len(s) - 1
	i >= 0
	{
		a := p * (int(s[i]) - 55)

		if s[i] > 47 && s[i] < 58 {
			a = p * int(s[i]-48)
		}
		r += a

		p *= n
	i--
	}

	Print(r)
}