package main
import . "fmt"
func main() {
	s := ""
	u := 0
	m := 100

	Scan(&s)
	for m < 1e3 {
		t := Sprint(m)
		i := 0
		for _, c := range s {
			if t[i] == byte(c) {
				i++
				if i == len(t) {
					u++
					break
				}
			}
		}
		m++
	}

	Print(u)
}