package main
import . "fmt"
func main() {
	x := 0
	s := "error"

	Scan(&x)
	if x > 0 && x < 10 {
		s = "nine"
	}

	Print(s)
}