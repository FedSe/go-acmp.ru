package main
import . "fmt"
func main(){
	var a, b, c, d int
	s:="NO"
	Scan(&a, &b, &c, &d)

	a-=c
	b-=d

	if a < 0 {
		a = -a
	}

	if b < 0 {
		b = -b
	}

	if a == b {
		s = "YES"
	}

	Print(s)
}