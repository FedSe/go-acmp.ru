package main
import . "fmt"
func main() {
	n:=0
	Scan(&n)
	var c[71]int
	c[0] = 1

	for
	i := 1
	i <= n
	{
		j := 0
		if i > 3 {
			j = i-3
		}

		for j < i {
			c[i] += c[j]
		j++
		}
	i++
	}

	Print(c[n])
}