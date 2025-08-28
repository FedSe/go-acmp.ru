package main
import . "fmt"
func main() {
	var a, b, c int
	Scan(&a, &b, &c)

	A := a + c - b
	C := a + b - c
	b += c - a

	s := C
	if b < A && b < C {
		s = b
	}
	if A < b && A < C {
		s = A
	}

	Print(s)
}