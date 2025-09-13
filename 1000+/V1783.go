package main
import . "fmt"
func main() {
	x := 0

	Scan(&x)
	x -= 9
	if x%9 == 0 {
		Print(x / 9)
		return
	}

	Print("error")
}