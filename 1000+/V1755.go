package main
import . "fmt"
func main() {
	h := 0
	m := 0
	s := "night"

	Scanf("%d:%d", &h, &m)
	m += h * 60
	if m > 269 && m < 690 {
		s = "morning"
	}
	if m > 689 && m < 930 {
		s = "day"
	}
	if m > 929 && m < 1410 {
		s = "evening"
	}

	Print(s)
}