package main
import . "fmt"
func main() {
	i := 0
	a := ""
	s := a

	for i < 3 {
		Scan(&s)
		a += s
		i++
	}

	s = "Draw"
	for _, l := range [][]int{
		{0, 1, 2}, {3, 4, 5}, {6, 7, 8},
		{0, 3, 6}, {1, 4, 7}, {2, 5, 8},
		{0, 4, 8}, {2, 4, 6}} {
		u := 0
		for _, j := range l {
			u += int(a[j])
		}
		if u == 264 {
			s = "Win"
		}
		if u == 237 {
			s = "Lose"
		}
	}

	Print(s)
}