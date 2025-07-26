package main
import . "fmt"
func main(){
	var a, b, c, i int
	s:="DRAW"

	for i < 4 {
		Scan(&a, &b)
		c+= a-b
	i++
	}

	if c>0 { s="1" } 
	if c<0 { s="2" }

	Print(s)
}