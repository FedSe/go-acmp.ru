package main
import . "fmt"
func main() {
	var a, m, l int
	s:=""
	Scan(&s)
 
	for _, r := range s {
		if r > 49 { a++ } else { a-- }
		if a < m { m = a }
		if a > l { l = a }
	}

	Print(l - m + 1)
}