package main
import . "fmt"
func main(){
	var c, b, a int
	Scan(&a, &b, &c)
	b -= a
	Print(a+b*c-b)
}