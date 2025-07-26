package main
import . "fmt"
func main() {
	var (
		f = []int{2, 3}
		s, n, b int
		k = 2
		i = 5
		p = 1
		a = 1228
	)

	for i < 10000 {
		j := 0
		for j < k {
			if f[j]*f[j] > i {
				break
			}
			if i%f[j] < 1 {
				goto A
			}
		j++
		}

		f = append(f, i)
		k++
A:
		if i == p*6-1 {
			i = p*6 + 1
			p++
		} else {
			i = p*6 - 1
		}
	}

	Scan(&n)
	for f[b] <= n/2 {
		for f[b]+f[a] > n {
			a--
		}
		if f[b]+f[a] == n {
			s++
		}
		b++
	}

	Print(s)
}