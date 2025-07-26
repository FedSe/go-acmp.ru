package main
import . "fmt"

func isLetter(n string) int {
	t:=0
	if (n >= "a" && n <= "z") || (n >= "A" && n <= "Z") {
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
		c+=isLetter(s)
	i++
	}

	Print(c)

}