package main
import . "fmt"
func main() {
	n := 0
	Scan(&n)
	Print([]any{"Mon", "Tues", "Wednes", "Thurs", "Fri", "Satur", "Sun"}[n-1], "day")
}