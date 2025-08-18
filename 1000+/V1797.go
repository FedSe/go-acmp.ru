package main
import . "fmt"
func main() {
	s := ""
	r := 1
	p := 0

	Scan(&s)
	for _, v := range s {
		d := int(v - 48)
		if d < 0 {
			p = d
		}
		if d > 0 {
			if p == -5 {
				r += d
			}
			if p == -3 {
				r -= d
			}
			if p == -6 {
				r *= d
			}
		}
	}

	Print(r)
}