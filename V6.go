package main
import . "fmt"
func main() {
	s := ""
	Scan(&s)
	if len(s) == 5 && s[2] == 45 &&
		s[0] > 64 && s[0] < 73 &&
		s[3] > 64 && s[3] < 73 &&
		s[1] > 48 && s[1] < 57 &&
		s[4] > 48 && s[4] < 57 {
		z := s[0] - s[3]
		c := s[1] - s[4]
		s = "NO"
		if z*z+c*c == 5 {
			s = "YES"
		}
	} else {
		s = "ERROR"
	}
	Print(s)
}