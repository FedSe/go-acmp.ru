package main
import (
	. "fmt"
	. "sort"
	. "os"
	b "bufio"
)
func main() {
	var (
		a [4]string
		i = 0
		m = b.NewScanner(Stdin)
	)

	for m.Scan() {
		a[i] = m.Text()
	i++
	}
	Strings(a[1:])

	Print(a[0], ": ", a[1], ", ", a[2], ", ", a[3])
}