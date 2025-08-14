package main
import . "fmt"
func main() {
	var (
		f = [1e4]int{2, 3}
		v = 0
		k = 2
		i = 5
		p = 1
	)

	Scan(&v)
	for v > 1 {
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

		f[k] = i
		k++
		j = 1
		for j < k {
			if k == f[j] {
				v--
				break
			}
			j++
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