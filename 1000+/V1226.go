package main
import . "fmt"
func main() {
	c := 0
	i := 0

	for i < 3 {
		s := ""
		Scan(&s)
		w := s[0]
		if w > 64 && w < 91 || w > 96 && w < 123 {
			c++
		}
		i++
	}

	Print(c)
}