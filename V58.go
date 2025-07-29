package main
import . "fmt"
func main() {
	var (
		a       [100][100]int
		k, n, m int
		S       = Scan
	)

	S(&k)
	for 0 < k {
		s := "YES"
		S(&n, &m)
		i := 0
		for i < n {
			j := 0
			for j < m {
				S(&a[i][j])
				j++
			}
			i++
		}
		i = 0
		for i < n-1 {
			j := 0
			for j < m-1 {
				if (a[i][j]+a[i+1][j]+a[i][j+1]+a[i+1][j+1])%4 < 1 {
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