package main
import . "fmt"
func main() {
	var n, a, c, i int
	b:=1
	Scan(&n)
	for c < n {
		a = b
		b = c
		c = a + b
		i++
	}
	if c == n {
		Print(1,"\n", i)
	} else {
		Print(0)
	}
}