package main
import . "fmt"
func main() {
	var (
		p                [100]int
		n, m, b, a, y, i int
		S                = Scan
	)

	S(&n)
	for i < n {
		S(&p[i])
		i++
	}

	S(&b, &m)
	for 0 < m {
		var t, s, j int
		for j < n {
			S(&i)
			s += i
			t += i * p[j]
			j++
		}
		t -= y
		if s == n {
			t += b
		}
		if t > a {
			a = t
		}
		Println(a)
		y += 2
		m--
	}
}