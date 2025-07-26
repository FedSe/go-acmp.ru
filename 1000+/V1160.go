package main
import (
	. "fmt"
	b "bufio"
	. "os"
)
func main() {
	var (
		v [1000000]int
		s = ""
		i = 1
		w = b.NewWriter(Stdout)
	)

	Scan(&s)
	n := len(s)
	Print("0 ")
	for i < n {
		j := v[i-1]
		for j > 0 && s[i] != s[j] {
			j = v[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		v[i] = j
		w.WriteString(Sprint(j, " "))
	i++
	}

	w.Flush()
}