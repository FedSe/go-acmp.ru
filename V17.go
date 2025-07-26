package main
import . "fmt"
func main() {
	var (
		n, l, k, i int
		p [30001]int
		a [30000]int
	)
	Scan(&n)

	for k < n {
		Scan(&a[k])
	k++
	}

	for i < n-1 {
	i++
		for {
			if a[l] == a[i] {
				l++
				break
			}
			if l < 1 {
				break
			}
			l = p[l]
		}
	p[i+1] = l
	}

	for {
		i = n - l
		if (n-1)%i < 1 {
			break
		}

		l = p[l]
	}

	Print(i)
}