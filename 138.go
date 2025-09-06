package main
import . "fmt"
func main() {
	var (
		r             [1e4][3]int
		d             [101]int
		n, m, i, j, l int
		f             = 1
	)

	Scan(&n, &m)
	for i < n {
		d[i] = 3e4
		i++
	}

	d[0] = 0
	for j < m {
		i = 0
		for i < 3 {
			Scan(&r[j][i])
			if i < 2 {
				r[j][i]--
			}
			i++
		}
		j++
	}

	for f > 0 {
		f = 0
		k := 0
		for k < m {
			i = r[k][0]
			j = r[k][1]
			o := d[i] + r[k][2]
			if d[i] < 3e4 && d[j] > o {
				f = 1
				d[j] = o
			}
			k++
		}
	}

	for l < n {
		Println(d[l])
		l++
	}
}