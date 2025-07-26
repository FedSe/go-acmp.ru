package main
import . "fmt"
func main() {
    var (
        w [18]int
        n, t, i, j int
    )
    Scan(&n)

    for i < n {
        Scan(&w[i])
    i++
    }

    for _, e := range w[:n] {
        t += e
    }

    a := t
    for j < (1 << n) {
        s := 0
        i = 0
        for i < n {
            if (j >> i) & 1 > 0 {
                s += w[i]
            }
        i++
        }
        d := t - 2*s
        if d < a && d > -1 {
            a = d
        }
    j++
    }

    Print(a)
}