package main
import . "fmt"
func main() {
	var (
		n, j, i, s, l int
		a             [52]int
	)

	Scan(&n)
	for l < n {
		Scan(&a[l])
		l++
	}

	for j < 10 {
		f := a[i]
		i++
		g := a[i]
		h := a[i+1] + 10
		s += h + g
		if f != 10 {
			s += f - h
			if f+g == 10 {
				s += h - f - g
			}
			i++
		}
		j++
	}

	Print(s)
}