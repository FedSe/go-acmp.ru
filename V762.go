package main
import . "fmt"
func main() {
	var n, t, k int
    
	Scan(&n)
	t = 1
	for t < n {
		t *= 3
		k++
	}
    
	Print(k)
}