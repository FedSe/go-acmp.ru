package main
import . "fmt"
func main(){
	s:="NO"
	var a, b, c, d int
	Scan(&a, &b, &c, &d)
	a-=c
	b-=d
	if a < 0 {a = -a}
	if b < 0 {b = -b}

	if (a == 1 && b == 2) || (a == 2 && b == 1) {
		s = "YES"
	}

	Print(s)
}