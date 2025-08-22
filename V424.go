package main
import . "fmt"
func main() {
	n := 0
	t := 0
	w := 1

	Scan(&n)
	for w < n {
		w *= []int{9, 2}[t%2]
		t++
	}

	Print([]any{"Ollie", "Stan"}[t%2], " wins.")
}