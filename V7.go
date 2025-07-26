package main
import (
	. "fmt"
	. "math/big"
)
func main() {
	var a, b, c Int

	Scan(&a, &b, &c)

	if b.Cmp(&a) > 0 {
		a = b
	}
	if c.Cmp(&a) > 0 {
		a = c
	}

	Print(&a)
}