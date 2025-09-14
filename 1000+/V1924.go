package main
import . "fmt"
func main() {
	var N, M, X, Y int
	Scan(&N, &M, &X, &Y)
	Print(N*(N+1)*M*(M+1)/4 - X*(N-X+1)*Y*(M-Y+1))
}