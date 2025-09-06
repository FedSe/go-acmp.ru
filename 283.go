package main
import . "fmt"
func main() {
	a := "Yes"
	s := a
	c := 0

	Scan(&s)
	s += "A"
	for _, v := range s[1:] {
		if 64 < v && v < 91 {
			if c < 1 || c > 3 || s[0] > 90 {
				a = "No"
			}
			c = -1
		}
		c++
	}

	Print(a)
}