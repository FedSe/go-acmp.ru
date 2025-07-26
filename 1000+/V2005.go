package main
import . "fmt"
func main() {
	a:=0
	Scan(&a)
	a+=a%2
	a/=2
	Print(a*a)
}