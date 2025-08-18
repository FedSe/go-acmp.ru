package main
import . "fmt"
func main() {
	a := "Romario"
	s := a

	Scan(&s)
	n := len(s)

	if n > 1 && s[n-1]&1 > 0 && s[n-2]&1 > 0 || n&1 < 1 {
		a = "Vito"
	}

	Print(a)
}