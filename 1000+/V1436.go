package main
import . "fmt"
func main() {
	n := 0
	i := 1
	P := Print

	Scan(&n)
	for i <= n {
		j := 0
		for j < (n-i)/2 {
			P(".")
			j++
		}
		j = 0
		for j < i {
			P("A")
			j++
		}
		P(`
`)
		i += 2
	}
}