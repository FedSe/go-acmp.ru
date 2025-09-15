package main
import (
	. "fmt"
	. "math/big"
)
func main() {
	var (
		n, j int
		d    [10]Int
		t    Int
		m    = map[int][]int{
			0: {4, 6},
			1: {6, 8},
			2: {7, 9},
			3: {4, 8},
			4: {0, 3, 9},
			6: {0, 1, 7},
			7: {2, 6},
			8: {1, 3},
			9: {2, 4}}
	)

	Scan(&n)
	for j < 9 {
		j++
		if j != 8 {
			d[j].SetInt64(1)
		}
	}

	for 1 < n {
		n--
		var p [10]Int
		j = 0
		for j < 10 {
			for _, v := range m[j] {
				p[v].Add(&p[v], &d[j])
			}
			j++
		}
		d = p
	}

	for _, v := range d {
		t.Add(&t, &v)
	}

	Print(&t)
}