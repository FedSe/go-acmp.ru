package main
import . "fmt"
func main() {
	var (
		a       [101][101]int
		d       [101]int
		n, i, f int
		S       = Scan
	)

	S(&n)
	for i < n {
		d[i] = 999
		j := 0
		for j < n {
			S(&a[i][j])
			j++
		}
		i++
	}

	S(&i, &f)

	i--
	f--
	d[i] = 0
	q := []int{i}

	for len(q) > 0 {
		i = q[0]
		q = q[1:]
		j := 0
		for j < n {
			if a[i][j] > 0 && d[j] > d[i]+1 {
				d[j] = d[i] + 1
				q = append(q, j)
			}
			j++
		}
	}

	i = d[f]
	if i == 999 {
		i = -1
	}
	Print(i)
}