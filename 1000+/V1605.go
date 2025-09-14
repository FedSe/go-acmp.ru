package main
import . "fmt"
func main() {
	x := 0
	y := 0
	Scan(&x, &y)
	Print(x*(x-1)/2, 2*x*y+y*(y-1))
}