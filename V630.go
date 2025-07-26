package main
import (
    b "bufio"
    . "fmt"
    . "os"
)
func main() {
    g := 0
    r := b.NewReader(Stdin)
    Fscan(r, &g)
    
    for g > 0 {
        var (
            C, D, E [10001]int
            n, a, v, w, t, z, j int
            o = int(^uint(0) >> 1)
            s = "Accepted"
        )
        Fscan(r, &n)
        for j < n {
            Fscan(r, &a, &v)
            C[a]++
            C[v]--
            D[a] += j
            D[v] -= j
            j++
        }
    
        for z < 10000 {
            t += C[z]
            if t < o {
                o = t
            }
            w += D[z]
            if t == 1 {
                E[w] = 1
            }
            z++
        }
    
        for 0 < n {
        n--
            if E[n] * o != 1 {
                s = "Wrong Answer"
            }
        }
        Println(s)
    g--
    }
}