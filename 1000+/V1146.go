package main
import . "fmt"
func main(){
	c:=' '
	s:="No"
	Scan(&c)

	if c < 10 {
		s="Yes"
	}

	Print(s)
}