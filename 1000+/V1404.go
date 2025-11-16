package main
import . "fmt"
func main() {
	s := ""
	Scan(&s)
	for _, c := range s {
		if c > 121 {
			c = 96
		}
		Printf("%c", c+1)
	}
}