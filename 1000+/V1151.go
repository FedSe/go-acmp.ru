package main
import . "fmt"
func main() {
	var l, u, d int
	a := "No"
	s := a

	Scan(&s)
	for _, t := range s {
		if t > 96 {
			l = 1
		}
		if t > 64 && t < 91 {
			u = 1
		}
		if t < 58 {
			d = 1
		}
	}

	if l+u+d > 2 && len(s) > 11 {
		a = "Yes"
	}

	Print(a)
}