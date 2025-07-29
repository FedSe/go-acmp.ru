package main
import . "fmt"
func main() {
	var (
		a          [1e5]int
		j, i, n, k int
		S = Scan
	)

	S(&n)
	for i < n {
		S(&a[i])
		i++
	}

	S(&k)
	if k > 0 {
		k %= n
	} else {
		k = (k%n + n) % n
	}

	i = n - k
	for i < n {
		Println(a[i])
		i++
	}

	for j < n-k {
		Println(a[j])
		j++
	}
}