package main
import ."fmt"
func main() {
	a := [10000]int{2, 3, 4, 7, 13}
	b := [10000]int{1, 5, 6, 8, 9, 10, 11, 12}
	n:=0
	Scan(&n)

	j := 8
	for
	i := 5
	i < n
	{
		a[i] = b[i-1] + b[i-3]
		for
		k := a[i-1] + 1
		k < a[i] && j < n
		{
			b[j] = k
		k++
		j++
		}
	i++
	}

	n--
	Print(a[n], "\n", b[n])
}