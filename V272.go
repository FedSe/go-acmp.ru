package main
import (
	b "bufio"
	."fmt"
	. "os"
)
func main() {
	var y, z, i, a int
	Scan(&z, &y)

	s := b.NewScanner(Stdin)
	s.Split(b.ScanWords)

	for s.Scan() {
		Sscan(s.Text(), &a)
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