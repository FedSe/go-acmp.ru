package main
import (
	. "fmt"
	. "sort"
)
func main() {
	s := ""

	Scan(&s)

	b := []byte(s)
	Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})

	for {
		Println(string(b))
		n := b
		f := len(n)
		i := f - 2
		for i >= 0 && n[i] >= n[i+1] {
			i--
		}
		if i < 0 {
			break
		}
		f--
		for f >= 0 && n[f] <= n[i] {
			f--
		}
		n[i], n[f] = n[f], n[i]

		n = n[i+1:]
		i = 0
		j := len(n)
		for i < j {
			j--
			n[i], n[j] = n[j], n[i]
			i++
		}
	}
}