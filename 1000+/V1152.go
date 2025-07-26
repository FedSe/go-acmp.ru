package main
import . "fmt"
func main() {
	s := ""
	f := ' '
	Scan(&s)
	for _, n := range s {
		Printf("%c%c",f, n)
		f = 35
	}
}