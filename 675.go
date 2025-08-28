package main
import . "fmt"
func main() {
	n:=0
	m:=0
	s:=""
	Scan(&n, &m)

	for 0 < n {
		Scan(&s)
		k := 0
		for _, r := range s {
			if r < 47 { k++ }	
		}

		if k < m {
			m = k
		}
	n--
	}

	Print(m)
}