package main
import . "fmt"
func main() {
	t := ""
	var s []rune

	for {
		_, e := Scan(&t)
		if e != nil {
			break
		}

		a := 0
		for _, c := range t {
			if c == 40 || c == 91 || c == '{' {
				s = append(s, c)
			} else if len(s) > 0 && ((c == 41 && s[len(s)-1] == 40) || (c == 93 && s[len(s)-1] == 91) || (c == '}' && s[len(s)-1] == '{')) {
				s = s[:len(s)-1]
			} else {
				a = 1
				break
			}
		}

		if len(s) > 0 {
			a = 1
			s = s[:0]
		}

		Print(a)
	}
}