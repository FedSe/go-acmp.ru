package main
import . "fmt"
var (
	b, p, a [1e4]int
	k, n, i int
	f       = 1 > 0
)

func F(j int) {
	b[j] = 1
	k++
	i := 0
	for f && i < n {
		if a[j+i*100] > 0 {
			if b[i] < 1 {
				p[i] = j
				F(i)
			} else {
				f = p[j] == i
			}
		}
		i++
	}
}

func main() {
	s := "NO"
	Scan(&n)
	for i < n {
		j := 0
		for j < n {
			Scan(&a[i+j*100])
			j++
		}
		i++
	}

	F(0)
	if f && k == n {
		s = "YES"
	}

	Print(s)
}