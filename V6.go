package main
import . "fmt"
func main() {
	a := "NO"
	s := a
	Scan(&s)
	z := int(s[0])
	if len(s) < 5 || s[2] != 45 ||
		z < 65 || z > 72 ||
		s[3] < 65 || s[3] > 72 ||
		s[1] < 49 || s[1] > 56 ||
		s[4] < 49 || s[4] > 56 {
		a = "ERROR"
	} else {
		z = (z-int(s[3]))*(int(s[1])-int(s[4]))
		if z == 2 || z == -2 { a = "YES" }
	}
	Print(a)
}