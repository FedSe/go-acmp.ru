package main
import . "fmt"
func main() {
	n := 0
	s := ""
	Scan(&n, &s)

	L := 1
	R := n
	for _, c := range s {
		n = L
		if c > 61 {
			n = R
			R--
			L--
		}
		Println(n)
		L++
	}

	Print(R)
}