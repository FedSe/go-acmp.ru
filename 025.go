package main
import . "fmt"
func main(){
	a:=1
	b:=1
	s:="="
	Scan(&a,&b)

	if a > b {
		s=">"
	}
	if a < b {
		s="<"
	} 

	Print(s)
}