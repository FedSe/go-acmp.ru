package main
import . "fmt"
func main() {
	var K, N, C int
	Scan(&K, &N, &C)
	Print(K * (N - C))
}