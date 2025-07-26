package main
import (
   . "fmt"
   . "regexp"
)

func main() {
	n := 1
	Scan(&n)
	for n > 0 {
		s := "Yes"
		p := s
		Scan(&p)

		if !MustCompile("^[ABCEHKMOPTXY]{1}[0-9]{3}[ABCEHKMOPTXY]{2}$").MatchString(p) {
			s = "No"
		}

		Println(s)
	n--
	}
}