package main
import . "fmt"
func main() {
	var k, m, n int
	Scan(&k)
	for 0 < k {
		Scan(&n, &m)
		Println(19*m + (n+239)*(n+366)/2)
		k--
	}
}