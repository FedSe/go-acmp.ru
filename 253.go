package main
import . "fmt"
func main() {
	var s, b, e, f, c int
	Scan(&s, &b, &e, &f)
	
	b += s * 60
	f += e * 60
	if b > f {
		f += 1440
	}

	for b < f {
		s = b / 60
		for s > 12 {
			s -= 12
		}

		e = b % 60
		if e < 1 {
			c += s
		}
		if e == 30 {
			c++
		}
		b++
	}

	Print(c)
}