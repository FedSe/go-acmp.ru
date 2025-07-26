package main
import . "fmt"
func main(){
	s:=""
	a:=0
	i:=0
	Scan(&s)

	for i < len(s) - 4 {
		b:=s[i:i+5]
		if b == ">>-->" || b == "<--<<" {
			a++
		}
	i++
	}

	Print(a)
}