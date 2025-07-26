package main
import . "fmt"
func F(n int) int {
    u := 0
    for n > 0 {
        u += n % 10
        n /= 10
    }
    return u
}

func main() {
    v := 0

    for {
        _, e := Scan(&v)
        if e != nil {
            break
        }
        var (
            a []int
            d = F(v)
            w = v
            i = 3
        )

        for w%2 < 1 {
            a = append(a, 2)
            w /= 2
        }
        
        for i*i <= w {
            for w%i < 1 {
                a = append(a, i)
                w /= i
            }
        i += 2
        }
        
        if w > 2 {
            a = append(a, w)
        }
        
        i = 0
        if a[0] != v {
            v = 0
            for _, o := range a {
                v += F(o)
            }
            if d == v {
                i = 1
            }
        }
        Print(i)
        
    }
}