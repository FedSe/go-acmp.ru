package main
import . "fmt"
func main() {
	var (
		a [104][104]int
		d = []int{0, 1, 0, -1, 1, 0, -1, 0}
		s = ""
		m, r, i int
		x = 51
		y = x
		w = -1
	)
	a[x][y] = 1

	Scan(&s)
	for i < len(s) {
		if s[i] == 82 {
			r += 5
			r %= 4
		}
		if s[i] < 77 {
			r += 3
			r %= 4
		}
		if s[i] > 82 {
			x += d[r]
			y += d[r+4]
		m++
			if a[x][y] == 1 {
				w = m
				break
			}
			a[x][y] = 1
		}
	i++
	}

	Print(w)
}