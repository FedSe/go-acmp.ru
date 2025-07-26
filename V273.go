package main
import . "fmt"
func main() {
	s := ""
	Scan(&s)
	u := 0
	for
	m := 100
	m < 1000
	{
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