package main
import . "fmt"
func main() {
	var (
		a          [100001]int
		j, i, n, k int
	)

	Scan(&n)
	for i < n {
		Scan(&a[i])
		i++
	}

	Scan(&k)
	if k > 0 {
		k %= n
	} else {
		k = n - ((-k) % n)
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