package main
import . "fmt"
func main(){
	s:=1
	t:=1
	Scan(&s, &t)

	if t < s { t += 12 }

	Print(t-s)
}