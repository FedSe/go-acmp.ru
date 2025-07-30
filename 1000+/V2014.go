package main
import . "fmt"
func main() {
	a := 0
	b := 0
	Scan(&a, &b)
	Printf("%.1f", float64(a*b)/2.)
}