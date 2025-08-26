package main
import . "fmt"
func main() {
	var (
		s       []int
		n, k, d int
		S       = Scan
	)

	S(&n)
	for n > 0 {
		S(&k)
		if k < 3 {
			S(&d)
			if k < 2 {
				s = append([]int{d}, s...)
			} else {
				s = append(s, d)
			}
		} else {
			if k < 4 {
				d = s[0]
				s = s[1:]
			} else {
				k = len(s) - 1
				d = s[k]
				s = s[:k]
			}
			Println(d)
		}
		n--
	}
}