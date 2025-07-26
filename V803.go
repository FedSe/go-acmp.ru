package main
import . "fmt"
func main() {
	var a, b, k int
	Scan(&a, &b, &k)

	for 0 < k {
		a %= b
		a *= 10
	k--
	}

	Print(a / b)
}