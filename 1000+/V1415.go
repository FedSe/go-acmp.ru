package main
import . "fmt"
func main() {
	var k, a, b int

	Scan(&k, &a, &b)
	if a+b == 0 {
		k = 0
	}

	Print(k + 1)
}