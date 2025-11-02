package main
import . "fmt"
var s string
func F(l, r int) int {
	b := int(s[l] - 48)
	if l != r {
		m := l + 2
		a := 0
		for m < r {
			n := s[m]
			if n == 44 && a == 0 {
				break
			}
			if n < 41 {
				a++
			}
			if n == 41 {
				a--
			}
			m++
		}
		a = F(l+2, m-1)
		b = F(m+1, r-1)
		if a < b {
			a, b = b, a
		}
		if s[l] == 77 {
			b = a
		}
	}
	return b
}

func main() {
	Scan(&s)
	Print(F(0, len(s)-1))
}