package main
import . "fmt"
func main() {
	a := 0
	b := 0
	Scan(&a, &b)
	n := b
	if n < 0 { n = -n }
	Print((a%b + n) % n)
}