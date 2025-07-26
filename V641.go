package main
import . "fmt"
func main() {
	n := ""
	a := n
	i := 0
	Scan(&n)

	for i < len(n) {
		for
		j := i + 1
		j < len(n)
		{
			m := n[:i] + n[i+1:j] + n[j+1:]
			if m > a {
				a = m
			}
		j++
		}
	i++
	}

	Print(a)
}