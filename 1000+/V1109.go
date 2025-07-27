package main
import . "fmt"
func main() {
	a := 1
	Scan(&a)
	Printf(`The next number for the number %d is %d.
The previous number for the number %d is %d.`, a, a+1, a, a-1)
}