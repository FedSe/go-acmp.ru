package main
import . "fmt"
func main(){
	var l, u, d bool
	a:="No"
	s:=a

	Scan(&s)
	for _,t := range s {
		if t >= 'a' && t <= 'z' {
			l = true
		}
		if t >= 'A' && t <= 'Z' {
			u = true
		}
		if t >= '0' && t <= '9' {
			d = true
		}

	}

	if l && u && d && len(s) > 11 { a = "Yes"}
	Print(a)

}