package main
import . "fmt"
func main() {
	s := 0
	a := "Darkoff"

	Scan(&s)

	v := s * 3 / 100.
	b := 2 * (s / 50)
	d := 5 * (s / 100)
	if b >= d {
		a = "Beta"
	}
	if v >= b && v >= d {
		a = "Vendel"
	}

	Print(a, " Bank")
}