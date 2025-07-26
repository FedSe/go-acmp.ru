package main
import . "fmt"
func main(){
	c:=' '
	Scanf("%c", &c)
	if c >= 'a' && c <= 'z' {
		c += 'A' - 'a'
	}
	Printf("%c", c)
}