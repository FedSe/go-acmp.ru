package main
import . "fmt"
func main() {
	c := 0
	s := ""
	Scan(&s)

	for len(s) > 1 {
		n := 0
		for _, i := range s {
			n += int(i-48)
		}
		c++
		s = Sprint(n)	
	}

	Println(s, c)

}