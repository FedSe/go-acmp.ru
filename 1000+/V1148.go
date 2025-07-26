package main
import . "fmt"
func main(){
	c:=""
	Scan(&c)
	a:=c[0]
	if c >= "a" && c <= "z" {
		a = c[0] + 'A' - 'a'
	}
	if c >= "A" && c <= "Z" {
		a = c[0] + 'a' - 'A'
	}

	Printf("%c", a)
}