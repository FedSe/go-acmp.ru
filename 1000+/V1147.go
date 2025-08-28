package main
import . "fmt"
func main() {
	c := ' '
	Scanf("%c", &c)
	if c > 96 && c < 123 {
		c -= 32
	}
	Printf("%c", c)
}