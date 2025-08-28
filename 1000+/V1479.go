package main
import . "fmt"
func main() {
	a := 0
	b := 0

	Scan(&a, &b)
	if b < 0 {
		b = -b
	}

	Print((a%b + b) % b)
}