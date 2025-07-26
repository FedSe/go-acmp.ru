package main
import . "fmt"
func main() {
	var a, b, d, e, f, i int
	Scan(&a, &b)

	for i < a {
	i++
		for 
		j := 0
		j < b
		{
		j++
			if i*j%5 < 1 {
				f++
			} else if i*j%3 < 1 {
				e++
			} else if i*j%2 < 1 {
				d++
			}

		}
	}

	Printf("RED : %d\nGREEN : %d\nBLUE : %d\nBLACK : %d", d, e, f, a*b-d-e-f)
}