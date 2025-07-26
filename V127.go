package main
import . "fmt"
func main() {
	var (
		a [101][101]int
		d [101]int
		n, i, f int
	)

	Scan(&n)

	for i < n {
		d[i] = 999
		for
		j := 0
		j < n
		{
			Scan(&a[i][j])
		j++
		}
	i++
	}

	Scan(&i, &f)

	i--
	f--
	d[i] = 0
	q := []int{i}

	for len(q) > 0 {
		i = q[0]
		q = q[1:]

		for
		j := 0
		j < n
		{
			if a[i][j] > 0 && d[j] > d[i]+1 {
				d[j] = d[i] + 1
				q = append(q, j)
			}
		j++
		}
	}

	if d[f] == 999 {
		d[f] = -1
	}
	Print(d[f])

}