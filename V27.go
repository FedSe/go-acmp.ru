package main
import . "fmt"
func main() {
	var (
		a [100][100]int
		w, h, n, i, y, z, v, s int
	)

	Scan(&w, &h, &n)

	for 0 < n {
		Scan(&i, &y, &z, &v)
		for i < z {
			for
			j := y
			j < v
			{
				if a[i][j] < 1 {
					s++
					a[i][j] = 1
				}
			j++
			}
		i++
		}
	n--
	}

	Println(w*h - s)
}