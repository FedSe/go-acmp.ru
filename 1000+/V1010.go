package main
import . "fmt"
func main() {
	var (
		n, i, l, k int
		m          [26][26][26]int
		P          = Printf
	)

	Scan(&n)
	for l < n {
		i = 0
		for i < n {
			m[0][l][i] = 90 - l
			q := 1
			for q < n {
				m[q][i][l] = 90 - l
				q++
			}
			i++
		}
		l++
	}

	for k < n {
		l = 0
		for l < n {
			i = 0
			for i < n {
				P("%c", m[k][l][i])
				i++
			}
			P(" ")
			l++
		}
		P(`
`)
		k++
	}
}