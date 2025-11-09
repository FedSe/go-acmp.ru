package main
import . "fmt"
func main() {
	var (
		h, t, i int
		a       = "BREAK"
		s       = []int{510, 595, 610, 695, 710, 795, 825, 910, 925, 1010, 1025, 1110}
		P       = Print
	)

	Scanf("%d:%d", &h, &t)
	t += h * 60
	for i < 12 {
		if t >= s[i] && t < s[i+1] {
			P(i/2 + 1)
			return
		}
		i += 2
	}

	if t >= s[11] {
		a = "EVENING"
	}
	if t < s[0] {
		a = "MORNING"
	}

	P(a)
}