package main
import . "fmt"
func main() {
	a := 0
	Scan(&a)
	Print(a/2*4 + (a-1)/2*2 + 1)
}