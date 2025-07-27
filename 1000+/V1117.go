package main
import . "fmt"
func main() {
	n := 0
	Scan(&n)

	k := n - 1
	r := 540 + 45*n + 20*(k/2) + 5*(k%2)

	Println(r/60, r%60)
}