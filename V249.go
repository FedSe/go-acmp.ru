package main
import . "fmt"
func main() {
	var c [100][100]int
	s := ""
	h := 0
	Scan(&s)

	if len(s) < 1 {
		Print(0)
		return
	}

	n := len(s)

	for i := range c {
		for j := range c[i] {
			c[i][j] = 101
		}
	}

	for h <= n {
	h++
		for
		f := 0
		f+h-1 < n
		{
			r := f + h - 1
			w := &c[f][r]
			m := s[r] - s[f]
			v := m < 3 && m > 0
			if h < 3 {
				*w = h
				if v { *w = 0 }
			} else {
				b := c[f+1][r-1]
				if v && *w > b { *w = b }
				for
				a := f
				a < r
				{
					b = c[f][a]+c[a+1][r]
					if *w > b {
						*w = b
					}
				a++
				}
			}
		f++
		}
	}

	Print(c[0][n-1])
}