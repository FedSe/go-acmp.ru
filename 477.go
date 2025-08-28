package main
import . "fmt"
func main() {
	s := ""
	a := 0
	P := Print

	Scan(&s)
	for _, c := range s {
		a += int(c - 48)
	}

	a %= 3
	if a < 1 {
		P(2)
	} else {
		P(1, a)
	}
}