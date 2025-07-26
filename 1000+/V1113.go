package main
import . "fmt"
func main(){
	a:=0
	b:=0
	Scan(&a, &b)
	c:=2
	if a % b < 1 || b % a < 1 {
		c = 1
	}

	Print(c)
}