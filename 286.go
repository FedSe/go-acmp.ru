package main
import (
	. "fmt"
	. "math/big"
)
func main() {
	var a, b Rat
	s := "="

	Scan(&a, &b)
	f := a.Cmp(&b)
	if f > 0 {
		s = ">"
	}
	if f < 0 {
		s = "<"
	}

	Print(s)
}