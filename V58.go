package main
import . "fmt"
func main() {
	var (
		a [100][100]int
		k, n, m int
	)
	Scan(&k)

	for 0 < k {
		s := "YES"
		Scan(&n, &m)
		
		for
		i := 0
		i < n
		{
			for
			j := 0
			j < m
			{
				Scan(&a[i][j])
			j++
			}
		i++
		}

		for
		i := 0
		i < n-1
		{
			for
			j := 0
			j < m-1
			{
				if (a[i][j] + a[i+1][j] + a[i][j+1] + a[i+1][j+1]) % 4 < 1 {
					s = "NO"
				}
			j++
			}
		i++
		}

		Println(s)
	k--
	}
}