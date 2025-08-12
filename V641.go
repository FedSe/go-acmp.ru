package main
import . "fmt"
func main() {
	n := ""
	a := n
	i := 0

	Scan(&n)
	k := len(n)
	for i < k {
		i++
		j := i
		for j < k {
			m := n[:i-1] + n[i:j] + n[j+1:]
			if m > a {
				a = m
			}
			j++
		}
	}

	Print(a)
}