package main
import . "fmt"
func main() {
	p := 1
	s := ""
	Scan(&s)

	for _, c := range s {
		if c < 66 && p < 3 {
			p = 3 - p
		}
		if c == 66 && p > 1 {
			p = 5 - p
		}
		if c > 66 && p != 2 {
			p = 4 - p
		}
	}

	Print(p)
}