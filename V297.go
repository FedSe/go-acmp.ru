package main
import . "fmt"
func main() {
	c := 0
	s := ""

	Scan(&s)
	for _, n := range s {
		c += int("1000001021"[n-48] - 48)
	}

	Print(c)
}