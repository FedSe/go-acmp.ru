package main
import . "fmt"
func main() {
	c := 0
	i := 0

	for i < 3 {
		s := ""
		Scan(&s)
		if s[0] > 47 && s[0] < 58 {
			c++
		}
		i++
	}

	Print(c)
}