package main
import . "fmt"
func main(){
	s:="NO"
	n:=s
	Scan(&n)

	if n[0] + n[1] + n[2] == n[3] + n[4] + n[5] {
		s="YES"
	}

	Print(s)
}