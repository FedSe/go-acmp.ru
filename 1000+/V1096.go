package main
import (
	b "bufio"
	. "fmt"
	. "os"
	. "strings"
)
func main() {
	r := b.NewScanner(Stdin)
	r.Scan()
	n := len(Fields(r.Text())) - 1
	r.Scan()
	w := Fields(r.Text())
	Print(w[n%len(w)])
}