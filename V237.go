package main
import (
	. "fmt"
	. "sort"
)
func main() {
	var (
		T [10][10]byte
		n, m, j, i int
		s = ""
		v []byte
	)

	Scan(&n, &m)

	for j < n {
		Scan(&s)
		for
		x := 0
		x < len(s)
		{
			T[j][x] = s[x]
		x++
		}
	j++
	}
	for 0 < m {
		Scan(&s)
		for
		x := 0
		x < len(s)
		{
			for
			i := 0
			i < n
			{
				for
				j = 0
				j < n
				{
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
		for
		j = 0
		j < n
		{
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

	for _, z := range v {
		Print(string(z))
	}
}