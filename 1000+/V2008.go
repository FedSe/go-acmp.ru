package main
import . "fmt"
func main() {
	i := 0
	Scan(&i)
	j := i / 10 * 2
	if i%10 > 0 {
		j++
	}
	Print(j)
}