package main
import . "fmt"
func main() {
	n := 0
	x := 0
    
	Scan(&n)
	for 0 < n {
		n--
		Println(x)
		x += 32768 + n*7
	}
}