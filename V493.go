package main
import (. "fmt"
. "strings"
)
var (n, m, k, i int
	a []string
)

func A(){
	a = append(a, "." + Repeat(".", m) + ".")
}

func main() {
	Scan(&n, &m)

	A()
	for i < n {
		t:=""
		Scan(&t)
		a = append(a, "." + t + ".")
	i++
	}
	A()

	for
	i := 1
	i < n+1
	{
		for
		j := 1
		j < m+1
		{
			if a[i-1][j] + a[i][j+1] + a[i+1][j] + a[i][j-1] + a[i][j] == 230 {
				k++
			}
		j++
		}
	i++
	}

	Print(k)
}