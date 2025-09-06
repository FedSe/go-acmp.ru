package main
import . "fmt"
func main() {
	a := 1
	s := ""

	Scan(&s)
	for _, r := range s {
		if r > 48 {
			Printf("%c", 96+a)
			a = 0
		}
		a++
	}
}