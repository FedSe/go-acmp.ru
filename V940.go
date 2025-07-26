package main
import . "fmt"
func main() {
	a:=1
	s:=""
	Scan(&a, &s)
	Print(s[:a-1] + s[a:])
}