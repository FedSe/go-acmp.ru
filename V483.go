package main
import (
	. "fmt"
	. "strconv"
)

func main() {
	n := 0
	o := 0
	s := ""
	Scan(&n, &s)

	for 0 < n {
	n--
		if s[n] > 48 {
			o++
		}
	}

	for _, q := range FormatInt(int64(o), 2) {
		Print(q-48)
	}
}