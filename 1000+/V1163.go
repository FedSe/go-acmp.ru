package main
import . "fmt"
func main() {
	n := 0
	a := 0.
	b := a

	Scan(&n, &a, &b)
	Print(int(float64(int(float64(n)*(1+a/100)+.501))*(1-b/100) + .501))
}