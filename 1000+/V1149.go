package main
import . "fmt"
func main(){
	s:=""
	c:=0
	Scan(&s)
	for _,t := range s {
		if t < 49 {
			c++
		}
	}

	Print(c)
}