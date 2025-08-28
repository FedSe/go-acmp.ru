package main
import . "fmt"
func main() {
	c := [71]int{1}
	n := 0
	i := 0

	Scan(&n)
	for i < n {
		i++
		j := 0
		if i > 3 {
			j = i - 3
		}
		for j < i {
			c[i] += c[j]
			j++
		}
	}

	Print(c[n])
}