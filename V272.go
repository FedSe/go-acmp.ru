package main
import (
	b "bufio"
	. "fmt"
	. "os"
)
func main() {
	var y, z, i, a int
	Scan(&z, &y)

	s := b.NewReader(Stdin)
	for {
		_, e := Fscan(s, &a)
		if e != nil {
			break
		}

		i++

		if i%2 < 1 {
			if a > y {
				y = a
			}
		} else {
			if a < z {
				z = a
			}
		}
	}
	Print(y + z)
}