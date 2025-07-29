package main
import (
	d "bufio"
	. "fmt"
	. "os"
	. "strconv"
)
func main() {
	var (
		a, b, c             [2e3]int
		n, m, i, k, j, l, h int
		s                   = d.NewScanner(Stdin)
		S                   = Scan
		P                   = Print
	)
	s.Split(d.ScanWords)

	S(&n)
	for k < n {
		S(&a[k])
		k++
	}
	if n < 3 {
		i = a[0]
		if n > 1 {
			h = a[0]
			if h > a[1] {
				h = a[1]
			}
			i += a[1]
			for l < n {
				S(&a[l])
				l++
			}
			if h > a[0] {
				h = a[0]
			}
			if h > a[1] {
				h = a[1]
			}
			i += a[0] + a[1]
		}
		P(i - h)
		return
	}

	x := a[0] + a[1] + a[2]
	for h < n-2 {
		m = a[h] + a[h+1] + a[h+2]
		if m > x {
			x = m
		}
		h++
	}
	for j < n {
		S(&b[j])
		j++
	}
	for l < n-2 {
		m = b[l] + b[l+1] + b[l+2]
		if m > x {
			x = m
		}
		l++
	}
	for i < n-1 {
		m = b[i] + b[i+1] + a[i+1]
		if m > x {
			x = m
		}
		m = a[i] + b[i] + b[i+1]
		if m > x {
			x = m
		}
		m = a[i] + b[i+1] + a[i+1]
		if m > x {
			x = m
		}
		m = b[i] + a[i] + a[i+1]
		if m > x {
			x = m
		}
		i++
	}

	j = 2
	for j < n {
		i = 0
		for i < n && s.Scan() {
			c[i], _ = Atoi(s.Text())
			i++
		}
		i = 0
		for i < n {
			m = a[i] + b[i] + c[i]
			if m > x {
				x = m
			}
			i++
		}
		i = 0
		for i < n-2 {
			m = c[i] + c[i+1] + c[i+2]
			if m > x {
				x = m
			}
			i++
		}
		i = 0
		for i < n-1 {
			m = c[i] + c[i+1] + b[i+1]
			if m > x {
				x = m
			}
			m = b[i] + c[i] + c[i+1]
			if m > x {
				x = m
			}
			m = b[i] + c[i+1] + b[i+1]
			if m > x {
				x = m
			}
			m = c[i] + b[i] + b[i+1]
			if m > x {
				x = m
			}
			i++
		}
		i = 0
		for i < n {
			a[i] = b[i]
			b[i] = c[i]
			i++
		}
		j++
	}
	P(x)
}