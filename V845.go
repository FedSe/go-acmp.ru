package main
import . "fmt"
func main() {
	x := ""
	y := x
	r := 0

	Scan(&x, &y)
	d := int(x[len(x)-1] - 48)
	for i := range y {
		r = (r*10 + int(y[i])) % 4
	}

	Print("0000111124863971464655556666793184269191"[d*4+(r+3)%4] - 48)
}