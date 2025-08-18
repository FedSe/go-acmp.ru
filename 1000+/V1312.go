package main
import (
	. "fmt"
	. "math/big"
)
func main() {
	var (
		n int64
		N = NewInt
		w = N(2)
		s = ""
		b = 0
	)

	Scan(&s)
	for len(s) >= b {
		s = s[b:]
		n++
		b = len(N(1).Exp(w, N(n), nil).String())
	}

	Print(n - 1)
}