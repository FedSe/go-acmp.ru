package main
import . "fmt"
func main() {
	a := 0
	b := 0
	n := 10
    
	Scan(&a, &b)

	if a < 0 || b < 0 {
		n = 1
	}
    
	if b >= 0 && b != 2 || a >= 0 && a != 4 {
		n = 0
	}

	Print(n)
}