package main
import . "fmt"
func main() {
	n:=0
	b:=0
	Scan(&n)
	b = n / 10
	if n % 10 > 0 {
		b++
	}
	Print(b)
}