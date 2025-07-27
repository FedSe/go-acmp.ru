package main
import . "fmt"
func main() {
	n := 0
	w := 0

	Scan(&n)
	for n%3 > 0 {
		n -= 2
		w++
	}

	if w > 0 {
		Println(2, w)
	}

	if n/3 > 0 {
		Print(3, n/3)
	}
}