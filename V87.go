package main
import . "fmt"
func main() {
	var (
		k, i int
		s = ""
		v [240]string
		m = make(map[string]int)
	)

	Scan(&s)
	for s != "ENDOFINPUT" {
		v[i] = s
		i++
		m[s] = 1
		Scan(&s)
	}

	for _, t := range v {
		for
		i = 1
		i < len(t)
		{
			if m[t[:i]] * m[t[i:]] > 0 {
				k++
				break
			}
		i++
		}
	}

	Print(k)
}