package main
import . "fmt"
func main() {
	var i, n, m int

	Scan(&n)
	for i < n {
		i++
		Scan(&m)
		if m < 438 {
			Print("Crash ", i)
			return
		}

	}

	Print("No crash")
}