package main
import . "fmt"
func main() {
	n := 0
	i := 1
	d := [101]int{1}

	Scan(&n)
	for i <= n {
		j := n
		for j >= i {
			d[j] += d[j-i]
			j--
		}
		i++
	}

	Print(d[n])
}