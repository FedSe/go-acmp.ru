package main
import . "fmt"
func main() {
	var (
		s []int
		n, k, d int
	)

	Scan(&n)

	for n > 0 {
		Scan(&k)

		if k < 3 {
			Scan(&d)
			if k == 1 {
				s = append([]int{d}, s...)
			} else {
				s = append(s, d)
			}
		} else {
			if k < 4 {
				d = s[0]
				s = s[1:]
			} else {
				d = s[len(s)-1]
				s = s[:len(s)-1]
			}
			Println(d)
		}
	n--
	}
}