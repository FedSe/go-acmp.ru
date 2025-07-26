package main
import . "fmt"
func main() {
	var (
		a [1000002]int
		m, n, c int
		i = 2
		j = 2
		M = 1000001
	)

	for i < M {
		a[i] = 1
	i++
	}
    
	for j <= M/2 {
		for
		i = j + j
		i <= M
		i += j {
			a[i] += j
		}
	j++
	}
    
	Scan(&m, &n)
    
	for
	i = m
	i <= n
	{
		j = a[i]
		if j >= m && j <= n && a[j] == i && i < j {
		c++
			Println(i, j)
		}
	i++
	}
    
	if c < 1 {
		Print("Absent")
	}
}