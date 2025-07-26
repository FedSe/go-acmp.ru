package main
import . "fmt"
func main() {
	var n, m, k, c int
	Scan(&n, &m, &k)

	c = m / k
	k--
	if k > m { k = m }

	k+=c
	Print(k * n)
	
}