package main
import . "fmt"
func main() {
	a := [1e4]int{2, 3, 4, 7, 13}
	b := [1e4]int{1, 5, 6, 8, 9, 10, 11, 12}
	n := 0
	j := 8
	i := 5

	Scan(&n)
	for i < n {
		a[i] = b[i-1] + b[i-3]
		k := a[i-1] + 1
		for k < a[i] && j < n {
			b[j] = k
			k++
			j++
		}
		i++
	}

	n--
	Println(a[n], b[n])
}