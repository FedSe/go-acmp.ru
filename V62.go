package main
import . "fmt"
func main(){
	s:="WHITE"
	i:=s
	Scan(&i)
	if i[0] % 2 == i[1] % 2 { s = "BLACK" }
	Print(s)
}