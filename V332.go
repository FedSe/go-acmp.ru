package main
import . "fmt"
func main() {
	var (
		c, b [63001]int
		n, i, j int
	)

	Scan(&n)

	for i < n {
		for
		j = i
		j < n
		{
		j++
			Scan(&c[i*251+j])
		}
	i++
		b[i] = 5000

		for
		j = 0
		j < i
		{
			x := b[j]+c[j*251+i]
			if x < b[i] {
				b[i] = x
			}
		j++
		}
	}

	Print(b[n])
}