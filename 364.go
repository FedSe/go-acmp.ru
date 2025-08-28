package main
import . "fmt"
func main() {
	var m, n, f int
	Scan(&m, &n)

	for _, v := range []int{
		6,
		28,
		496,
		8128,
		33550336,
		8589869056,
		137438691328,
		2305843008139952128} {
		if v >= m && v <= n {
			Println(v)
			f = 1
		}
	}

	if f < 1 {
		Print("Absent")
	}
}