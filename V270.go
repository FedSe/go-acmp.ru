package main
import (
	. "fmt"
	. "strings"
)
func main() {
	s := ""
	i := 0
	C := Contains

	Scan(&s)
	if s[0] < 96 || C(s, "__") || s[len(s)-1] == 95 {
		s = "0_"
	}

	a := C(s, "_")
	if a {
		for i < len(s) {
			if s[i] < 91 {
				Print("Error!")
				return
			}
			i++
		}
	}

	i = 0
	for i < len(s) {
		if a && s[i] == 95 {
			s = s[:i] + string(s[i+1]-32) + s[i+2:]
			i--
		}

		if !a && s[i] < 91 {
			s = s[:i] + "_" + ToLower(s[i:i+1]) + s[i+1:]
			i++
		}
		i++
	}

	Print(s)
}