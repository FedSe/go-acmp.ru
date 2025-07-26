package main
import . "fmt"
func main() {
	n:=0
	Scan(&n)
	i:=2

	for n%i > 0 {
		i++
	}

	Print(i)
}