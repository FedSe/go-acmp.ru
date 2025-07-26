package main
import . "fmt"
func main() {
	var (
		f [21]int
		x, y, b int
	)
	Scan(&x, &y)

	f[1] = 1
	f[2] = 1

	z := 2
	for z < x {
	z++
		f[z] = f[z-2] + f[z-1]
	}

	z = f[x-2]
	for (y-f[x-1]*b) % z > 0 {
		b++
	} 

	Print((y - f[x-1]*b) / z, b)
}