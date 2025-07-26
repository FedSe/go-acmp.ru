package main
import . "fmt"
func main() {
	var (
		f = [9999]int{2, 3}
		v = 0
		k = 2
		i = 5
		p = 1
	)

	Scan(&v)

	for v > 1 {
		for
		j := 0
		j < k
		{
			if f[j]*f[j] > i {
				break
			}
			if i%f[j] < 1 {
				goto A
			}
		j++
		}

		f[k] = i
		k++
		for
		z := 1
		z < k
		{
			if k == f[z] {
				v--
				break
			}
		z++
		}
A:
		if i == p*6-1 {
			i = p*6 + 1
			p++
		} else {
			i = p*6 - 1
		}
	}

	Print(f[k-1])
}