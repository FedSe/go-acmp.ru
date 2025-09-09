package main
import . "fmt"
func main() {
	var (
		m          []int
		n, k, j, i int
		q          = [12]int{1}
	)

	Scan(&n, &k)
	w := make([]int, n)
	for j < n {
		if j > 0 {
			q[j] = q[j-1] * j
		}
		w[j] = j + 1
		j++
	}

	k--
	for i < n {
		j = q[n-i-1]
		x := k / j
		m = append(m, w[x])
		w = append(w[:x], w[x+1:]...)
		k %= j
		i++
	}

	for _, v := range m {
		Println(v)
	}
}