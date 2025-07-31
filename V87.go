package main
import . "fmt"
func main() {
	var (
		k, i int
		s    = ""
		v    [240]string
		m    = make(map[any]int)
	)

	Scan(&s)
	for s != "ENDOFINPUT" {
		v[i] = s
		i++
		m[s] = 1
		Scan(&s)
	}

	for _, t := range v {
		i = 1
		for i < len(t) {
			if m[t[:i]]*m[t[i:]] > 0 {
				k++
				break
			}
			i++
		}
	}

	Print(k)
}