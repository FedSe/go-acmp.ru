package main
import . "fmt"
func main() {
	s := ""

	Scan(&s)
	var (
		a       [50]int
		n       = len(s)
		g       = 1
		l, i, j int
	)

	for i < n {
		a[i] = g * int(s[i]-48)
		g = -g
		l += a[i]
		i++
	}

	n--
	l -= a[n]
	g = l
	for j < n {
		j++
		l -= a[n-j] + a[n-j+1]
		if g < l {
			g = l
		}
	}

	Print(g)
}