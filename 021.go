package main
import . "fmt"
func main() {
	var a, g, i int
	l:=99999
	for i < 3 {
		Scan(&a)
		if a > g { g = a }
		if a < l { l = a }
	i++
	}

	Print(g-l)
}