package main
import . "fmt"
func main() {
	var (
		a [26]int
		k = 0
		i = 0
		s = ""
		c = 65
	)

	Scan(&s)
	for c < 91 {
		t := -1
		i := 0
		for i < 26 {
			if s[i] >= byte(c) && (t < 0 || s[i] < s[t]) {
				t = i
			}
			i++
		}

		if t < 0 {
			Print("NO")
			return
		}

		a[k] = t + 1
		k++
		s = s[:t] + "." + s[t+1:]
		c++
	}

	Println("YES")
	for i < 26 {
		Println(a[i])
		i++
	}
}