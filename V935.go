package main
import . "fmt"
func main() {
	var a, b, c, d int
	s:="NO"
	Scan(&a, &b, &c, &d)
	if (a + b + c + d) & 1 < 1 {s="YES"}
	Print(s)
}