package main
import . "fmt"
func main() {
	var (
		n, m, k int
		d [10]int
	)
	Scan(&n, &m)

	for
	i:=1
	i<=n
	{
		for
		j:=1
		j<=m
		{
			f := i*j
			for f > 0 {
				d[f%10]++
				f/=10
			}
		j++
		}
	i++
	}

	for k < 10 {
		Println(d[k])
	k++
	}
}