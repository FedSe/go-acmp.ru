package main
import ."fmt"
func main() {
	var (
		p [100]int
		n, m, b, a, y, i int
	)

	Scan(&n)

	for i < n {
		Scan(&p[i])
	i++
	}

	Scan(&b, &m)

	for 0 < m {
		t := 0
		s := 0
		for
		j := 0
		j < n
		{
			Scan(&i)
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