package main
import . "fmt"
func main() {
	var (
		p, s [4e4]int
		n    = 0
		a    = 2
		i    = 2
		z    = 0
	)
    
	for i < 4e4 {
		if s[i] < 1 {
			p[z] = i
			z++
			j := i * i
			for j < 4e4 {
				s[j] = 1
				j += i
			}
		}
		i++
	}

	Scan(&n)
	for 1 < n {
		i = a
		q := 1
		for _, v := range p[:z] {
			if v*v > i {
				break
			}
			if i%v < 1 {
				q = v
				for i%v < 1 {
					i /= v
				}
			}
		}
		if i > 1 {
			q = i
		}
		a += q
		n--
	}

	Print(a)
}