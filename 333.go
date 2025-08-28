package main
import . "fmt"
func main() {
	var (
		u                   [10]int
		n                   = ""
		t, i, j, a, b, c, l int
		P                   = Println
	)

	for l < 3 {
		Scan(&n)
		a = 0
		for _, h := range n {
			a |= 1 << (h - 48)
		}
		if l < 2 {
			c = b
			b = a
		}
		l++
	}

	l = a & b & c
	for i < 10 {
		if l&(1<<i) > 0 {
			t++
			u[i]++
		}
		i++
	}

	P(t)
	for j < 10 {
		if u[j] > 0 {
			P(j)
		}
		j++
	}
}