package main
import . "fmt"
func main() {
	a:="NO"
	s:=a
	d:=s
	Scan(&s, &d)
	v := []rune(s + d)
	f := v[2] - v[0]
	if f < 0 { f = -f }

	if v[3] - v[1] >= f && (v[0] + v[1] + v[2] + v[3]) & 1 < 1 {
		a = "YES"
	}

	Print(a)
}