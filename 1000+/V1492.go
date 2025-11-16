package main
import (
	. "fmt"
	. "math/bits"
)
func main() {
	var n, k, m int
	Scan(&n, &k, &m)
	Print((n*Len(uint(k+9)) + 7) / 8 * m)
}