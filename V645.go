package main
import . "fmt"
func main() {
	var k, h, w int
	Scan(&k)

	m := 2 * k
	for
	i := 1
	i <= k
	{
		j := k / i
		l := i - j
		if l < 0 { l = -l }

		f := l + k - i * j
		if f < m {
			m = f
			h = i
			w = j
		}
	i++
	}

	Print(h, w)
}