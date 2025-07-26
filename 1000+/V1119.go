package main
import . "fmt"
func main(){
	var a, b, c, d int
	s:="NO"
	Scan(&a, &b, &c, &d) 
	
	if a == c || b == d {
		s="YES"
	}

	Print(s)
}