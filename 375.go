package main
import (
	. "fmt"
	. "strings"
)
func main() {
	var (
		b       Builder
		x, y, a int
		G       = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		s       = G
		u       = ""
		W = b.WriteByte
	)

	Scan(&x, &y, &s)

	if s == "0" {
		u = s
	}

	for s != "0" {
		z := 0
		for z < len(s) {
			for z < len(s) && a < y {
				a = a*x + IndexByte(G, s[z])
				z++
				if b.Len() > 0 && a < y && z < len(s) {
					W(48)
				}
			}
			W(G[a/y])
			a %= y
		}
		u = G[a:a+1] + u
		s = b.String()
		b.Reset()
		a = 0
	}
	Print(u)
}