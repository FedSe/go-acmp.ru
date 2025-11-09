package main
import . "fmt"
func main() {
	var a, n, i int

	Scan(&a, &n)
	v := []int{a}
	for i < n {
		i++
		j := 0
		for a > 0 {
			d := a % 10
			j += d * d
			a /= 10
		}
		a = j
		j = 0
		for j < len(v) {
			if v[j] == a {
				Print(v[j+(n-j)%(i-j)])
				return
			}
			j++
		}
		v = append(v, a)
	}

	Print(v[len(v)-1])
}