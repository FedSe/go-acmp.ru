package main
import . "fmt"
func main() {
	s:=.0
	x:=s
	i:=s
	Scan(&x)

	for s < x {
		s += 1 / (i+2)
	i++
	} 

	Print(i, " card(s)")

}