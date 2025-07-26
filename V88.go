package main
import . "fmt"
func main() {
	var (
		a [101][101]int
		n, l, i int
		s = "Incorrect"
	)

	Scan(&n)
	p:=n*n

	for l < p {
		for
		j := 0
		j < p
		{
			Scan(&a[l][j])
		j++
		}
	l++
	}

	l = 1
	for i < p {
		r := map[int]int{}
		c := map[int]int{}
		for
		j := 0
		j < p
		{
			w := a[i][j]
			if w < 1 || w > p {
				l = 0
			}
			r[w]++
			c[a[j][i]]++

			if i % n < 1 && j % n < 1 {
				s := map[int]int{}
				for 
				r := 0
				r < n
				{
					for
					v := 0
					v < n
					{
						s[a[i+r][j+v]]++
					v++
					}
				r++
				}
				if len(s) != p {
					l = 0
				}
			}

		j++
		}

		if len(r) != p || len(c) != p {
			l = 0
		}

	i++
	}


	if l > 0 {
		s = "Correct"
	}

	Print(s)

}