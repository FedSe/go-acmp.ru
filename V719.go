package main
import (
	j "bufio"
	. "fmt"
	. "math"
	. "math/big"
	. "os"
)

func F(n float64) int {
	if n > 1 {
		n = Floor(.5*Log(2*Pi*n)/Log(10)+n*Log10(n/E)) + 1
	}
	return int(n)
}

func main() {
	var (
		s = ""
		a = 0
		g = 1.
		h = 1000.
		k = j.NewReader(Stdin)
		b = map[any]int{"1": 0}
		N = NewInt
		f = N(1)
		n = 1
	)

	for n < 101 {
		f.Mul(f, N(int64(n)))
		b[f.String()] = n
		n++
	}

	Fscan(k, &s)
	c, o := b[s]
	if o {
		a = c
		goto A
	}

	c = len(s)
	for F(h) < c {
		h *= 2
	}

	for g <= h {
		m := Floor((g + h) / 2)
		n := m
		if n < 1 {
			n = 1
		}
		e := F(n)
		if e < c {
			g = m + 1
		} else if e > c {
			h = m - 1
		} else {
			a = int(n)
			break
		}
	}
A:
	Print(a)
}