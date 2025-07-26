package main
import . "fmt"
func main() {
	m:=1
	n:=1
	Scan(&m, &n)
	m--
	n--
	Print(m*n)
}