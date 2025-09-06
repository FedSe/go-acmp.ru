package main
import . "fmt"
func main() {
	s := 0
	t := 0

	Scan(&s, &t)
	if t < s {
		t += 12
	}

	Print(t - s)
}