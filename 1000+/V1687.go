package main
import . "fmt"
func main() {
	var x, y, z int
	Scan(&x, &y, &z)
	Print(x*1024 + z/y*100)
}