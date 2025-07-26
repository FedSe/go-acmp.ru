package main
import . "fmt"
func main() {
	var (
		n, m, p, i, j int
		s = "YES"
		e [100]string
	)

	Scan(&n, &m)
	for j < n {
		Scan(&e[j])
	j++
	}

	for i < n {
		for
		j = 0
		j < m
		{
			l := "."
			Scan(&p)
			if p > 3 {
				l += "R"
			}
			if p % 2 == 1 {
				l += "B"
			}
			if p == 2 || p == 3 || p == 6 || p == 7 {
				l += "G"
			}
				
				for
				f := 0
				f < len(l)
				{
					if []byte(e[i])[j] == l[f] {
						p=9
					}
				f++
				}
			if p != 9 { s="NO" }
		j++
		}
	i++
	}

	Print(s)
}