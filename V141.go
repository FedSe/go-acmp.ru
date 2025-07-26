package main
import . "fmt"
var (
    a [100][100]int
    b, p [100]int
    k, n, i int
    f = 1>0
)

func d(j int) {
	b[j] = 1
	k++
	for
	i := 0
	f && i < n
	{
		if a[j][i] > 0 {
			if b[i] < 1 {
				p[i] = j
				d(i)
			} else {
				f = p[j] == i
			}
		}
	i++
	}
}

func main() {
	s:="NO"
	Scan(&n)
	for i < n {
		for
		j := 0
		j < n
		{
			Scan(&a[i][j])
		j++
		}
	i++
	}

	d(0)
	if f && k == n {
		s="YES"
	}
	Print(s)
}