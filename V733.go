package main
import . "fmt"
func main() {
	var n, m, a int
	s := 1

	Scan(&n, &m)
	if n < 2 {
		Print(1, m)
		return
	}
	x := []int{m}

	Scan(&m)
	if n < 3 {
		Print(2, x[0], m)
		return
	}

	x = append(x, m)
	if x[1] < x[0] {
		m = x[0]
	}

	for 2 < n {
		Scan(&a)
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

	Println(len(x))
	for _, s := range x {
		Println(s)
	}
}