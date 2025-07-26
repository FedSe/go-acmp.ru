package main
import . "fmt"
func main() {
    var n, i, j int
    Scan(&n)

    p := make([]int, n)
    for i < n {
        Scan(&p[i])
    i++
    }

    i = n - 2
    for i >= 0 && p[i] > p[i+1] {
        i--
    }

    if i < 0 {
        for j < n {
        j++
            Println(j)
        }
        return
    }

    j = n - 1
    for p[j] <= p[i] {
        j--
    }

    p[i], p[j] = p[j], p[i]

    i++
    n--
    for i < n {
        p[i], p[n] = p[n], p[i]
        i++
        n--
    }

    for _, v := range p {
        Println(v)
    }
}