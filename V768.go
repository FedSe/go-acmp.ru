package main
import . "fmt"
var (
	x, y [100]int
	a, b, n, r, l, i, j, k int
	t bool
)
func F() () {
	t = 1<0
	if (x[j]-x[i])*(x[k]-x[i])+(y[j]-y[i])*(y[k]-y[i]) == 0 {
		a = x[j] - x[i] + x[k]
		b = y[j] - y[i] + y[k]
		t = 1>0
	}
	if (x[i]-x[j])*(x[k]-x[j])+(y[i]-y[j])*(y[k]-y[j]) == 0 {
		a = x[i] - x[j] + x[k]
		b = y[i] - y[j] + y[k]
		t = 1>0
	}
	if (x[j]-x[k])*(x[i]-x[k])+(y[j]-y[k])*(y[i]-y[k]) == 0 {
		a = x[j] - x[k] + x[i]
		b = y[j] - y[k] + y[i]
		t = 1>0
	}
}

func main() {
	Scan(&n)
	for l < n {
		Scan(&x[l], &y[l])
	l++
	}

	for i < n {
		j = i + 1
		for j < n {
			k = j + 1
			for k < n {
				F()
				if t {
					l = k + 1
					for l < n {
						if x[l] == a && y[l] == b {
							r++
						}
					l++
					}
				}
			k++
			}
		j++
		}
	i++
	}

	Print(r)
}