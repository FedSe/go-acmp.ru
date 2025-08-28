package main
import . "fmt"
func main() {
	var (
		i = 0
		A = "Yes"
		a [4]string
	)

	Scan(&a[0])
	for i < 3 {
		q := A
		Scan(&q)
		a[i+1] = q
		j := 0
		for j < 3 {
			if a[i][j]*3 == a[i][j+1]+q[j]+q[j+1] {
				A = "No"
			}
			j++
		}
		i++
	}

	Print(A)
}