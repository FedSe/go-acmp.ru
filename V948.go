package main
import . "fmt"
func main(){
	a:=0
	b:=0
	Scan(&a, &b)
	b--
	Print(1+b/a, 1+b%a)
}