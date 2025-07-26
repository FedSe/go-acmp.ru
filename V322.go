package main
import . "fmt"
func main() {
	b:=1
	p:=0
	s:=""
	Scan(&s)

	for p+b <= len(s) {
		a := p + b
		Print(s[a-1:a])
		p = b
		b = a
	}
}