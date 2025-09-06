package main
import . "fmt"
func main() {
	s := ""
	Scan(&s)
	Print(1)
	i := len(s) - 1
	for s[i] < 49 {
		Print(0)
		i--
	}
}