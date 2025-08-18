package main
import . "fmt"
func main() {
	var (
		q, v             ([101][101]int)
		F                func(y, x int)
		n, m, c, y, j, i int
	)

	Scan(&n, &m, &i)
	for 0 < i {
		Scan(&y, &j)
		q[y][j] = 1
		i--
	}

	F = func(y, x int) {
		if y*x < 1 || y > n || x > m || v[y][x]+q[y][x] > 0 {
			return
		}
		v[y][x] = 1
		F(y-1, x)
		F(y+1, x)
		F(y, x-1)
		F(y, x+1)
	}

	for i < n {
		i++
		j = 0
		for j < m {
			j++
			if v[i][j]+q[i][j] < 1 {
				c++
				F(i, j)
			}
		}
	}

	Print(c)
}