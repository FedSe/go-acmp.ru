package main
import (
	. "fmt"
	b "bufio"
	. "os"
)
func main() {
	var (
		a, w, c, l, e int
		v = "ERROR"
		r = b.NewScanner(Stdin)
		m = 1>0
		n = -1
	)
	r.Scan()
	s := r.Text()
	o := len(s)-6

	F := func(u *int) {
		n++
		o += 2
		m = s[n] == 45
		if m { n++ }
		l = n
		for n < o && s[n] > 47 && s[n] < 58 {
			*u = *u*10 + int(s[n]-48)
			n++
		}
		if m { *u = -*u }
		if n == l {
			e = 1
		}
	}


	if o < -1 {
		s += v
		e = 1
	}

	F(&a)
	d := s[n]

	F(&w)
	if s[n] != 61 {
		e = 1
	}

	F(&c)
	if n < o {
		e = 1
	}

	if e < 1 {
		if d == 45 {
			v = "NO"
			m = a-w == c
		}
		if d == 43 { 
			v = "NO"
			m = a+w == c
		}
		if d < 43 {
			v = "NO"
			m = a*w == c
		}
		if d == 47 {
			v = "NO"
			m = w != 0 && a/w == c && a%w == 0
		}
		if m { v = "YES" }
	}

	Print(v)
}