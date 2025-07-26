package main
import (. "fmt"
. "math"
)
func main(){
	var a, b, c, d float64
	Scan(&a, &b, &c, &d)
	c-=a
	d-=b
	Print(Sqrt(c*c + d*d))
}