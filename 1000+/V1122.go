package main
import . "fmt"
func main(){
	var a, b, c, d int
	Scan(&a, &b, &c, &d)
	s:="NO"
	a-=c
	b-=d
	if a < 0 { a = -a }
	if b < 0 { b = -b }

	if a < 2 && b < 2 {
		s = "YES"
	}

	Print(s)
}