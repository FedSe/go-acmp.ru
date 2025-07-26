package main
import . "fmt"
func main(){
	var a, b, c int
	s:="NO"
	Scan(&a, &b, &c)

	if a * b == c {
		s="YES"
	}

	Print(s)
}