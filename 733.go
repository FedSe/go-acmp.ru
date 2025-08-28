package main
import . "fmt"
func main() {
	var (
		n, m, a int
		s       = 1
		S       = Scan
		P       = Println
	)

	S(&n, &m)
	if n < 2 {
		P(1, m)
		return
	}
	x := []int{m}

	S(&m)
	if n < 3 {
		P(2, x[0], m)
		return
	}

	x = append(x, m)
	if x[1] < x[0] {
		m = x[0]
	}

	for 2 < n {
		S(&a)
		for a > x[s] && x[s] < m && s > 0 {
			x = x[:len(x)-1]
			s--
		}
		x = append(x, a)
		s++
		if x[s] > m {
			m = x[s]
		}
		n--
	}

	P(len(x))
	for _, s := range x {
		P(s)
	}
}