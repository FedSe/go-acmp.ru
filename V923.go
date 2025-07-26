package main
import . "fmt"
var c = make(map[int]int)
func f(x int) int {
	if x < 3 {
		return 0
	}
	if x == 3 {
		return 1
	}

	if _, o := c[x]; !o {
		c[x] = f(x/2) + f(x-x/2)
	}
	return c[x]
}

func main() {
	var a int
	Scan(&a)
	Print(f(a))
}