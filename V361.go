package main
import . "fmt"
func main() {
	s := ""
	Scan(&s)
	var (
		a [27][101]int
		v, t bool
		i = 0
		n = len(s)
		l = n
	)
	for i < n {
		for
		j := i + 1
		j <= n
		{
			a[s[i]-96][j]++
		j++
		}
	i++
	}

	for !t && l > 1 {
		l--
		for
		i = 0
		i < n-l
		{
			for
			j := i + 1
			j <= n-l
			{
				v = 1 > 0
				for
				c := 1
				c <= 26
				{
					v = v && ((a[c][i+l]-a[c][i]) == (a[c][j+l]-a[c][j]))
				c++
				}
				if v {
					t = 1 > 0
				}
			j++
			}
		i++
		}
	}

	if !t { l = 0 }
	Print(l)
}