package main
import . "fmt"
func main() {
	var a, b, c, d int
	i := -100

	Scan(&a, &b, &c, &d)
	for i < 101 {
		if i*(i*i*a+i*b+c) == -d {
			Println(i)
		}
		i++
	}
}