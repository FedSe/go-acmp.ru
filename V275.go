package main
import . "fmt"
func main() {
	n := 0
	Scan(&n)
	for 0 < n {
		o := "No"
		s := o
		x := 0

		Scan(&s)

		for len(s) % 3 > 0 {
			s = "0" + s
		}

		for len(s) > 0 {
			x += int((s[0]-48)*4 + (s[1]-48)*2 + s[2]-48)
			s = s[3:]
		}

		if x % 7 < 1 {
			o = "Yes"
		}

		Println(o)
	n--
	}
}