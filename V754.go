package main
import . "fmt"
func main() {
	var m, a, i int

	for i < 3 {
		Scan(&m)
		if m < 94 || m > 727 {
			Print("Error")
			return
		}

		if m > a {
			a = m
		}
	i++
	}

	Print(a)
}