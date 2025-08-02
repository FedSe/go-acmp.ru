package main
import . "fmt"
func main() {
	var (
		c, N int
		B    = 2
		a    = ""
		d    = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	)

	Scan(&N)
	for B < 37 {
		n := N
		s := ""
		for n > 0 {
			s = string(d[n%B]) + s
			n /= B
		}
		p := 1
		i := 0
		l := len(s)
		for i < l/2 {
			if s[i] != s[l-1-i] {
				p = 0
			}
			i++
		}
		if p > 0 {
			a += Sprintln(B)
			c++
		}
		B++
	}
	d = `unique
`
	if c < 1 {
		d = `none
`
	}
	if c > 1 {
		d = `multiple
`
	}
	Print(d + a)
}