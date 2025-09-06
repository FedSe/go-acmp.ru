package main
import . "fmt"
func main() {
	n := 1
	s := "Error"

	Scan(&n)
	if n < 12 {
		s = "Autumn"
	}
	if n < 9 {
		s = "Summer"
	}
	if n < 6 {
		s = "Spring"
	}
	if n < 3 || n == 12 {
		s = "Winter"
	}

	Print(s)
}