package main
import . "fmt"
func main() {
	var (
		n, x, R, i, j int
		m             [100][100]int
		y             = -1
		C             = 1
	)
	Scan(&n)

	for i < n*n {
		x += R
		y += C
		if 0 <= x*y && x < n && y < n && m[x][y] < 1 {
			i++
			m[x][y] = i
		} else {
			x -= R
			y -= C
			if C == 1 || C == -1 {
				R = C
				C = 0
			} else if R == 1 || R == -1 {
				C = -R
				R = 0
			}
		}
	}

	for j < n {
		i = 0
		for i < n {
			Println(m[j][i])
			i++
		}
		j++
	}
}