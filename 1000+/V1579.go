package main
import . "fmt"
func main() {
	s := "No"
	n := ' '

	Scanf("%c", &n)
	if n > 90 { n = n-32 }
	if n < 91 && n > 64 && n != 74 && n != 85 && n != 87 {
		s = "Yes"
	}

	Print(s)
}