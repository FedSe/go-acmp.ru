package main
import . "fmt"
func main() {
	n:=0
	i:=0
	a:="NO"
	s:=a
	Scan(&s)

	if len(s) % 2 > 0 {
		s += "0"
	}

	for i < len(s) {
		n += int(s[i])-int(s[i+1])
	i+=2
	} 

	if n % 11 == 0 { a = "YES" }

	Print(a)

}