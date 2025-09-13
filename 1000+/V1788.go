package main
import . "fmt"
func main() {
	n := 0
	a := 1
	i := 1

	Scan(&n)
	for i*i < n {
		i++
		j := 0
		for n%i < 1 {
			j++
			n /= i
		}
		j /= 2
		for 0 < j {
			a *= i * i
			j--
		}
	}

	Print(a)
}