package main
import . "fmt"
func main() {
	var a, b, n int

	Scan(&a, &b, &n)
	n *= a*100 + b

	Print(n/100, n%100)
}