package main
import . "fmt"
func main() {
	var k, i, X, Y float64
	p := [5]bool{}

	for i < 5 {
		Scan(&X, &Y)
		x:=.0
		for
		j := 0
		j < 5
		{
			a := x-X
			if a*a + Y*Y < 101 {
				p[j] = 1>0
				break
			}
		x+=25
		j++
		}
	i++
	}

	for _, n := range p {
		if n { k++ }
	}

	Print(k)
}