package main
import (
	b "bufio"
	. "fmt"
	. "os"
)
func main() {
	var s, n, x, y, w int
	r := b.NewReader(Stdin)

	Scan(&s, &n)
	for 0 < n {
		Fscan(r, &w)
		if x+w <= s {
			x += w
			y -= w
		}
		y += w
		n--
	}

	Print(x, y)
}