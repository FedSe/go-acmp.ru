package main
import . "fmt"
func main() {
	var (
		r          [100]int
		n, i, j, a int
		S          = Scan
	)

	S(&n, &i)
	for j < n {
		S(&r[j])
		j++
	}

	for 0 < i {
		o := 1
		j = 0
		for j < n {
			S(&a)
			if a > r[j] {
				o = 0
			}
			j++
		}
		Println(o)
		i--
	}
}