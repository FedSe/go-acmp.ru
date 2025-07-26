package main
import . "fmt"
func main() {
	var (
		n, m, d, z, i int
		b [500]string
	)

	Scan(&n, &m)
	for i < n {
		Scan(&b[i])
	i++
	}

	for z < n {
		for
		i = 0
		i < m
		{
			if b[z][i] < 36 {
				d++
				k := i
				for k < m && b[z][k] < 36 {
					b[z] = b[z][:k] + "." + b[z][k+1:]
					k++
				}
				for
				x := z + 1
				x < z + k- i
				{
					for
					j := i
					j < k
					{
						b[x] = b[x][:j] + "." + b[x][j+1:]
					j++
					}
				x++
				}
			}
		i++
		}
	z++
	}

	Print(d)
}