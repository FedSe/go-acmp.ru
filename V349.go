package main
import . "fmt"
func main() {
	var (
		p [1e6+1]int
		i = 2
		n = i
		m = i
	)
	Scan(&m, &n)
	n++

	for i < n {
		if p[i] > 0 {
			i++
		} else {
			j := i * i
			for j < n {
				p[j] = 1
				j += i
			}
			i++
		}
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