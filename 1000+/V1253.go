package main
import . "fmt"
func main() {
	a := 0
	b := 0

	Scan(&a, &b)
	k := a

	for k%b > 0 {
		k += a
	}

	Print(k)
}