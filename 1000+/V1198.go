package main
import . "fmt"
func main() {
	var n, b, c int
	Scan(&n, &b, &b, &c)
	for j, v := range []int{n - b - c, b, c} {
		i := 0
		for i < v {
			Printf("%c", []int{65, 99, 48}[j]+i%2)
			i++
		}
	}
}