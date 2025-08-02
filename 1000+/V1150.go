package main
import (
	. "fmt"
	. "strings"
)
func main() {
	t := ""
	R := ReplaceAll
	Scan(&t)

	Print(R(R(t, "4", ""), "7", ""))
}