package main
import . "fmt"
func main() {
	var (
		i = 0
		A = "Yes"
		a [4]string
	)

	Scan(&a[0])
	for i < 3  {
		Scan(&a[i+1])
		for
		j := 0
		j < 3
		{
			if a[i][j] * 3 == a[i][j+1] + a[i+1][j] + a[i+1][j+1] {
				A = "No"
			}
		j++
		}
	i++
	}

	Print(A)
}