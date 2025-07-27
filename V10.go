package main
import . "fmt"
func main() {
	var a, b, c, d int
	Scan(&a, &b, &c, &d)

	i := -100
	for i < 101 {
		if i*(i*i*a+i*b+c) == -d {
			Println(i)
		}
		i++
	}
}