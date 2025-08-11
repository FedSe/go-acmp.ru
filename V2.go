package main
import . "fmt"
func main() {
	n := 0
    
	Scan(&n)
	a := n * (1 + n) / 2
	if n < 1 {
		a = n + 1 - a
	}

	Print(a)
}