package main
import (
	. "fmt"
	. "sort"
)
func main() {
	var (
		T          [10][10]byte
		n, m, j, i int
		s          = ""
		v          []byte
		S          = Scan
	)

	S(&n, &m)
	for j < n {
		S(&s)
		x := 0
		for x < len(s) {
			T[j][x] = s[x]
			x++
		}
		j++
	}

	for 0 < m {
		S(&s)
		x := 0
		for x < len(s) {
			i := 0
			for i < n {
				j = 0
				for j < n {
					if T[i][j] == s[x] {
						T[i][j] = 3
						goto A
					}
					j++
				}
				i++
			}
		A:
			x++
		}
		m--
	}

	for i < n {
		j = 0
		for j < n {
			if T[i][j] != 3 {
				v = append(v, T[i][j])
			}
			j++
		}
		i++
	}

	Slice(v, func(i, j int) bool {
		return v[i] < v[j]
	})

	Print(string(v))
}