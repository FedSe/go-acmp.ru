package main
import . "fmt"
func main(){
	var a, b, c, i int
	Scan(&a, &b, &c)

	c*=a+b

	for c > 0 {
		c-=8
	i++
	}

	Print(i)
}