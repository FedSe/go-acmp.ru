package main
import (
	. "fmt"
	. "strings"
)
func main() {
	a := "error"
	s := a

	Scan(&s)
	s = ToLower(s)
	if s == "green" {
		a = "go"
	}
	if s == "red" {
		a = "wait"
	}

	Print(a)
}