package main
import . "fmt"
func main() {
	a := "NO"
	s := a

	Scan(&s)
	if s[0] == s[3] && s[1] == s[2] {
		a = "YES"
	}

	Print(a)
}