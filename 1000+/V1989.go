package main
import (
	. "fmt"
	. "strings"
)
func main() {
	w := "5555"
	s := w
	a := 0

	Scan(&s)
	for Contains(s, w) {
		s = Replace(s, w, "45", 1)
		a++
	}

	Print(a)
}