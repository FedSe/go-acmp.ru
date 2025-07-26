package main
import . "fmt"
func main(){
	a:=1
	b:=1
	s:=""
	Scan(&a, &b, &s)

	if s[0] < 98 || (s[2] < 98 && a < b) || (s[2] == 'e' && a > b) { a=b }

	Print(a)
}