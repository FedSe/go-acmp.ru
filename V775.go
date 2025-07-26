package main
import . "fmt"
func main() {
	n:=""
	m:=n
	Scan(&n, &m)
	if m == "0" {
		Print("NO")
	} else {
		l := len(n)-1
		Printf("%s%c", n[:l], n[l] + 1)
	}
}