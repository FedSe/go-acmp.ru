package main
import . "fmt"
func main() {
	var (
		q       [70][70]bool
		d, p    [71]int
		x       []string
		s       = ""
		P       = Println
		l, o, e int
		i       = 3
	)

	Scan(&s)
	n := len(s)
	for l < n {
		q[l][l] = 1 > 0
		if l < n-1 {
			q[l][l+1] = s[l] == s[l+1]
		}
		l++
	}

	for i <= n {
		l = 0
		for l <= n-i {
			j := l + i - 1
			q[l][j] = s[l] == s[j] && q[l+1][j-1]
			l++
		}
		i++
	}

	for o <= n {
		d[o] = o
		p[o] = o - 1
		i = 0
		for i < o {
			if q[i][o-1] && d[i]+1 < d[o] {
				d[o] = d[i] + 1
				p[o] = i
			}
			i++
		}
		o++
	}

	i = n
	for i > 0 {
		l = p[i]
		x = append(x, s[l:i])
		i = l
	}

	i = len(x)
	P(i)
	for e < i {
		i--
		x[e], x[i] = x[i], x[e]
		e++
	}

	for _, v := range x {
		P(v)
	}
}