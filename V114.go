package main
import . "fmt"
func main() {
    var n, k, b int
    Scan(&n, &k)

    k--
    a := k
    for n > 1 {
        a, b = (a + b)*k, a
        n--
    }

    Print(a + b)
}