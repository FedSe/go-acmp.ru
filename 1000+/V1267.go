package main
import . "fmt"
func main() {
	i := 0
	m := map[any]int{}
	for i < 5 {
		s := ""
		Scan(&s)
		m[s] = 1
		i++
	}

	Print(len(m))
}