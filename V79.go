package main
import . "fmt"
func main() {
	a:=0
	b:=0
	Scan(&a, &b)
	c := a % 10
	for 1 < b {
		a = a * c % 10
	b--
	}
	Print(a % 10)
}