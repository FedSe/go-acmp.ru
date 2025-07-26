package main
import . "fmt"
func main() {
	s := "0000111124863971464655556666793184269191"
	x := s
	y := s

	Scan(&x, &y)
	d := int(x[len(x)-1] - 48)
	r := 0

	for i := range y {
		r = (r*10 + int(y[i]-48)) % 4
	}

	Print(s[d*4+(r+3)%4] - 48)
}