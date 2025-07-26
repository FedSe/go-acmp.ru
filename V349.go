package main
import . "fmt"
func main() {
	var (
		p [1000000]int
		i=2
		n=i
		m=i
	)
	Scan(&m, &n)
	n++

	for i < n {
		if p[i] > 0 {
			i++
			continue
		}
		for
		j := i * i
		j < n
		{
			p[j] = 1
		j += i
		}
	i++
	}

	for m < n {
		if p[m] < 1 {
			Println(m)
			i = 0
		}
	m++
	}

	if i > 0 {
		Print("Absent")
	}
}