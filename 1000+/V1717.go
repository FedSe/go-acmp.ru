package main
import . "fmt"
func main() {
	x := 0
	a := 0

	for {
		_, q := Scan(&x)
		if q != nil {
			break
		}
		a += x
	}

	Print(a)
}