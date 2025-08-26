package main
import . "fmt"
func main() {
	var (
		d                [100][100]int
		N, M, l, j, i, q int
		r                = &q
		F                = func(w int) {
			w++
			if *r > w {
				*r = w
			}
		}
	)

	Scan(&N, &M)
	for l < N {
		j = 0
		for j < M {
			r = &d[l][j]
			Scan(&q)
			*r = 1e9 * (1 - q)
			if l > 0 {
				F(d[l-1][j])
			}
			if j > 0 {
				F(d[l][j-1])
			}
			j++
		}
		l++
	}

	j = N
	for j > 0 {
		j--
		l = M
		for l > 0 {
			l--
			r = &d[j][l]
			if N-1 > j {
				F(d[j+1][l])
			}
			if M-1 > l {
				F(d[j][l+1])
			}
		}
	}

	for i < N {
		l = 0
		for l < M {
			Println(d[i][l])
			l++
		}
		i++
	}
}