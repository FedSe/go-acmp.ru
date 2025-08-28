package main
import (
	. "fmt"
	. "strings"
)
func main() {
	var n, m, k, i, j int
	R := Repeat

	Scan(&n, &m)
	a := [103]string{R(".", m+2)}

	for j < n {
		t := ""
		Scan(&t)
		j++
		a[j] = "." + t + "."
	}
	a[j+1] = R(".", m+2)

	for i < n {
		i++
		j = 0
		for j < m {
			j++
			if a[i-1][j]+a[i][j+1]+a[i+1][j]+a[i][j-1]+a[i][j] == 230 {
				k++
			}
		}
	}

	Print(k)
}