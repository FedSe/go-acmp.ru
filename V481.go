package main
import . "fmt"
func main() {
	var d [30][30]int
	s := ""

	Scan(&s)
	n := len(s)

	i := n
	for i > 0 {
		i--
		d[i][i] = 1
		j := i + 1
		for j < n {
			d[i][j] = d[i+1][j] + d[i][j-1] + 1
			if s[i] != s[j] {
				d[i][j] -= d[i+1][j-1] + 1
			}
			j++
		}
	}

	Print(d[0][n-1])
}