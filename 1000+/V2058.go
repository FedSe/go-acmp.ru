package main
import . "fmt"
func main() {
    var n, m, k int
    Scan(&n, &m, &k)
    n--
    m--
    i := n/k
    if n > m {
        i = m/k
        n, m = m, n
    }
    n -= i * k
    m -= i * k + n

    if n > 0 {
        i++
    }
    
    i += m/k

    if m > m/k * k {
        i++
    }
 
    Print(i)
}