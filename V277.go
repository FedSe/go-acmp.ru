package main
import . "fmt"
func main() {
	var a, b, c, d, i int
	Scan(&a, &b, &c)

	if a != 0 {
		Print(a)
		d++
	}

	s := "x"
	for i < 2 {
		if b != 0 {
			if d * b > 0 {
				Print("+")
			}
			if b * b != 1 {
				Print(b)
			}
			if b == -1 {
				s = "-" + s
			}
			Print(s)
			d++
		}
	b = c
	s = "y"
	i++
	}


	if d < 1 {
		Print(0)
	}
}