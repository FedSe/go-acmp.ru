package main
import (
	b "bufio"
	. "fmt"
	. "os"
)
func main() {
	var n, x, y int
	r := b.NewReader(Stdin)

	Scan(&n, &x)
	for 1 < n {
		Fscan(r, &y)
		Println(x + y)
		x = y
		n--
	}

	Print(x)
}