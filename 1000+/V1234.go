package main
import . "fmt"
func main() {
	var (
		a [100][100]int
		n, i, j int
	)

	Scan(&n)
	for i < n {
		for
		j = 0
		j < n
		{
			Scan(&a[i][j])
		j++
		}
	i++
	}

	n--
	for
	c := n
	c >= 0
	{
		for
		j = n
		j >= 0 
		{
			Print(a[j][c], " ")
		j--
		}
	c--
	}
}