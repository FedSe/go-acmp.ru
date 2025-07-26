package main
import . "fmt"
func main() {
	var (
		n, k, i, l int
		a [50]string
	)
	Scan(&n)

	for i < n {
		Scan(&a[i])
		for
		j := 0
		j < n
		{
			if a[i][j] < 68 { k++ }
		j++
		}
	i++
	}

	k /= 2
	for l < n {
		for
		j := 0
		j < n
		{
			i = 2
			if k > 0 { i = 1 }
			Print(i)
			
			if a[l][j] < 68 { k-- }
		j++
		}
		Println()
	l++
	}
}