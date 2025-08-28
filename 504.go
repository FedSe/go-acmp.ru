package main
import . "fmt"
func main() {
	k := 0
	Scan(&k)
	Print([]any{"GCV", "VGC", "CVG"}[k%3])
}