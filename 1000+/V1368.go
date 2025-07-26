package main
import . "fmt"
func main() {
	var (
		a [51][51]int
		n, s, t, i, k int
		F = 99999999
	)

	Scan(&n, &s, &t)
	s--
	t--

	for i < n {
		for
		j := 0
		j < n
		{
			Scan(&a[i][j])
			if a[i][j] < 0 {
				a[i][j] = F
			}
		j++
		}
	i++
	}

	for k < n {
		for
		i = 0
		i < n
		{
			for
			j := 0
			j < n
			{
				f := a[i][k] + a[k][j]
				if a[i][j] > f {
					a[i][j] = f
				}
			j++
			}
		i++
		}
	k++
	}

	if a[s][t] < F {
		Print(a[s][t])
	} else {
		Print(-1)
	}
}