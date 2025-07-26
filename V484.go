package main
import . "fmt"

var e = []int{1, 1}
var s = []int{0, 0}
type O struct { l, w int }

func F(x int) int {
    if x > len(e) {
        F(x - 1)
    }
    if x == len(e) {
        e = append(e, e[x-1] + e[x-2])
    }
    return e[x]
}

func H(n int) int {
    if n > len(s) {
        H(n - 1)
    }
    if n == len(s) {
        s = append(s, F(n-1) + s[n-1])
    }
    return s[n]
}

func T(v O) []O {
    if v.l == 1 && v.w == 1 {
        return []O{{1, 1}}
    }
    if v.l == 2 && v.w == 2 {
        return []O{{2, 2}, {1, 1}}
    }
    if v.w <= F(v.l-1) {
        b := T(O{v.l - 1, v.w})
        for i := range b {
            b[i].l++
        }
        return append(b, O{1, 1})
    }
    r := T(O{v.l - 2, v.w - F(v.l-1)})
    for i := range r {
        r[i].l += 2
        r[i].w += F(v.l-1-i)
    }
    return append(r, O{2, 2}, O{1, 1})
}

func main() {
    var g int
    l := 1
    Scan(&g)
    for {
        if H(l+1) >= g {
                for _, o := range T(O{l, g - H(l)}) {
                Println(H(o.l) + o.w)
            }
            break
        }
    l++
    }
}