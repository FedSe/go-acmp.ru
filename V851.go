package main
import . "fmt"
func main() {
	b := 0
	g := 0
	s := ""

	Scan(&b, &g)
	for b > 0 || g > 0 {
		if b > 0 {
			s += "B"
			b--
		}
		if g > 0 {
			s += "G"
			g--
		}
		if g > 0 {
			s += "G"
			g--
		}
	}

	Print(s)
}