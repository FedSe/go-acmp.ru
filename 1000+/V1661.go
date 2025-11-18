package main
import . "fmt"
func main() {
	var p, k, t, r int
	Scan(&p, &k, &t, &r)
	p = (p + k - 1) / k
	Print(p*t + (p-1)*r)
}