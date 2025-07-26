package main
import . "fmt"
func main() {
	var a, i, j int
	Scan(&i, &j)
	b:= 1

	for i * j > 0 {
		if i < j {
			j %= i
		} else {
			i %= j
		}
	}

	for i+j > 0 {
		f := (a + b) % 1000000000
 		a = b
		b = f
	j--
	}

	Print(a)
}