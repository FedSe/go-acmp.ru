package main
import . "fmt"
func main() {
	var a, b, c int
	Scan(&a, &b, &c)
	Print((a*b+a*c+c*b)*2, a*b*c)
}