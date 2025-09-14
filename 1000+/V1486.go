package main
import . "fmt"
func main() {
	a := "Off"
	s := a

	Scan(&s)
	A := s[3] == 43
	if s[0] == 43 && (s[2] == 43 || A) || s[1] == 43 && A {
		a = "On"
	}

	Print(a)
}