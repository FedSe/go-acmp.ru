package main
import . "fmt"
func main() {
	var (
		s = ""
		p = s
		a = 10000.
		x, i, n, u, t float64
	)

	Scan(&p, &p)
	for i < 4 {
		Scan(&s, &s)
		if s != "out" {
			Scan(&s, &s)
			Sscan(s[5:], &t)
			if a > t {
				a = t
			}
			if x < t {
				x = t
			}
			u += t
		n++
		}
	i++
	}

	i = 4-n
	Print("Ping statistics for ", p+":\nPackets: Sent = 4 Received = ", n, " Lost = ", i, " (", i*25, "% loss)\n")
	if n > 0 {
		Print("Approximate round trip times:\nMinimum = ", a, " Maximum = ", x, " Average = ", int(u/n+0.5))
	}
}