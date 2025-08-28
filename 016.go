package main
import . "fmt"
func main() {
	n := 0
	i := 0
	d := [101]int{1}

	Scan(&n)
	for i < n {
		i++
		j := n
		for j >= i {
			d[j] += d[j-i]
			j--
		}
	}

	Print(d[n])
}