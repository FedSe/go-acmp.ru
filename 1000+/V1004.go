package main
import (
	. "fmt"
	. "strings"
)
func main() {
	var (
		q []string
		P = Print
		w = ""
		t = 1
	)

	for t > 0 {
		t, _ = Scan(&w)
		if HasSuffix(w, ".") {
			w = w[:len(w)-1]
			P(" ", w)
			i := len(q)
			for i > 0 {
				i--
				P(" ", q[i])
			}
			q = nil
			P(".")
		} else {
			q = append(q, w)
		}
	}
}