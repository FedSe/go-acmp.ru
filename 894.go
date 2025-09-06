package main
import (
	. "fmt"
	. "math"
)
func main() {
	r := .0
	s := r
	Scan(&s, &r)
	Print(Sqrt(r*r - s/Pi))
}