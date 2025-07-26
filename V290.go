package main
import . "fmt"
func main() {
	var (
		f, t [100]string
		n, m, a, o, i, j, x int
	)

	Scan(&n, &m)
	
	for i < n {
		Scan(&f[i])
	i++
	}
	Scan(&o, &x)

	for j < o {
		Scan(&t[j])
	j++
	}

	x -= m
	o -= n
	for 0 <= x {
		for
		y := 0
		y <= o
		{
			for
			i = x
			i < x+m
			{
				for
				j = y
				j < y+n
				{
					if f[j-y][i-x] - t[j][i] > 200 {
						goto A
					}
				j++
				}
			i++
			}
			a++
A:
		y++
		}
	x--
	}

	Print(a)
}