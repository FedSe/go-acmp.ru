package main
import . "fmt"
func main(){
	k:=1
	s:="VGC"
	Scan(&k)
	k %= 3
	if k < 1 { s="GCV" }
	if k > 1 { s="CVG" }

	Print(s)
}