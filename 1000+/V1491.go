package main
import . "fmt"
func main() {
	n := 0
	k := 0
	Scan(&n, &k, &k)
	Print(n * k / (k - 1))
}