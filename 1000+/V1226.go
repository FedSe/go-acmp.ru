package main
import . "fmt"
func main() {
	c := 0
	i := 0

	for i < 3 {
		s := ""
		Scan(&s)
		if s[0] > 64 && s[0] < 91 || s[0] > 96 && s[0] < 123 {
			c++
		}
		i++
	}

	Print(c)
}