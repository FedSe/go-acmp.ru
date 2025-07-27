package main
import . "fmt"
func main() {
	var n, a, i int
	b := 1
	Scan(&n)

	for a < n {
		a, b = b, a+b
		i++
	}

	if a == n {
		Println(1, i)
	} else {
		Print(0)
	}
}