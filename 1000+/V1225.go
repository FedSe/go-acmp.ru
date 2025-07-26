package main
import . "fmt"

func isDigit(n string) int {
	t:=0
	if n >= "0" && n <= "9" {
		t++
	}
	return t
}

func main(){
	i:=0
	c:=0
	s:=""

	for i < 3 {
		Scan(&s)
		c+=isDigit(s)
	i++
	}

	Print(c)

}