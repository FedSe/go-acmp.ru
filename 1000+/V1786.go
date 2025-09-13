package main
import . "fmt"
func main() {
	var a, b, c, d, x, y int
	s := ""
    
	Scan(&a, &b, &c, &d, &x, &y)
	e := a
	if x > a {
		e = x
	}
	if x > c {
		e = c
	}

	z := b
	if y > b {
		z = y
	}
	if y > d {
		z = d
	}

	if z == d {
		s = "N"
	}
	if z == b {
		s = "S"
	}
	if e == c {
		s += "E"
	}
	if e == a {
		s += "W"
	}

	Print(s)
}