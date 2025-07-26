package main
import . "fmt"
func main(){
	a:=0
	Scan(&a)

	if a < 2 { a = 0 }
	a/= 2 - a & 1
	
	Print(a)
}