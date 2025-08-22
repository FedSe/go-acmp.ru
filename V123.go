package main
import . "fmt"
func main() {
	d := [81][81]int{{1}}
	s := ""
	i := 0

	Scan(&s)
	n := len(s)
	for i < n {
		j := 0
		for j < n {
			v := d[i][j]
			m := s[i]
			w := m == 63
			if (m < 41 || w) && j < n {
				d[i+1][j+1] += v
			}
			if (m == 41 || w) && j > 0 {
				d[i+1][j-1] += v
			}
			j++
		}
		i++
	}

	Print(d[n][0])
}