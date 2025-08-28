package main
import . "fmt"
func main() {
	var (
		n, i, s int
		a[100][100]int
	)

	Scan(&n)
	for i < n {
		for
		j := 0
		j < n
		{
			Scan(&a[i][j])
		j++
		}
	i++
	}

	i = 3000
	for s < n {
		for
		v := s+1
		v < n
		{
			for
			w := v+1
			w < n
			{
				t := a[s][v] + a[v][w] + a[w][s]
				if t < i { i = t }
			w++
			}
		v++
		}
	s++
	}

	Print(i)
}