package main
import . "fmt"
func main() {
	c := 0
	s := ""

	Scan(&s)
	for _, n := range s {
		if n < 49 || n == 54 || n > 55 { c++ }
		if n == 56 { c++ }
	}

	Print(c)
}