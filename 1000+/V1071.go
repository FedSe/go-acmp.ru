package main
import . "fmt"
func main() {
	i := 0
	Scan(&i)

	a := i-1
	if i % 2 < 1 { a += 2 }

	Print(a)
}