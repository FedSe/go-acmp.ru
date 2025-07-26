package main
import . "fmt"
func main() {
	p:=1
	s:=""
	Scan(&s)

	for _, c := range s {
		if c < 66 {
			if p < 2  {p++} else if p== 2 {p--}
		}
		if c == 66 {
			if p == 2 {p++} else if p > 2 {p--}
		}
		if c > 66 {
			if p < 2  {p=3} else if p > 2 {p=1}
		}
	}
	Print(p)

}