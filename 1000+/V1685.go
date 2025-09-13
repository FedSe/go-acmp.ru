package main
import . "fmt"
func main() {
	n := 0
	m := 0
	a := "NO"

	Scan(&n, &m)
	n--
	m--
	if n/4 == m/4 {
		a = "YES"
	}
	if n&1 < 1 {
		a += `
LOW
`
	} else {
		a += `
HIGH
`
	}
	if m&1 < 1 {
		a += "LOW"
	} else {
		a += "HIGH"
	}

	Print(a)
}