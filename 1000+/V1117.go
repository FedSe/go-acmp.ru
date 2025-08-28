package main
import . "fmt"
func main() {
	n := 0
	Scan(&n)

	k := n - 1
	k = 540 + 45*n + k/2*20 + k%2*5

	Print(k/60, k%60)
}