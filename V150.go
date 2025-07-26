package main
import . "fmt"
func main() {
	var (
		n, s, a int
		F[100][100]int
		i = 1
	)
	Scan(&n, &s)
	n++

	for i < n {
		for
		j := 1
		j < n
		{
			Scan(&F[i][j])
		j++
		}
	i++
	}

	i = 1
	for i < n {
		for
		k := 1
		k < n
		{
			for
			j := 1
			j < n
			{
				if F[k][i] * F[i][j] > 0 {
					F[k][j] = 1
				}
			j++
			}
		k++
		}
	i++
	}

	for n > 1 {
	n--
		if n != s && F[s][n] > 0 {
			a++
		}
	}

	Print(a)
}