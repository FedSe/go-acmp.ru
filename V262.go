package main
import (
	. "container/heap"
	. "fmt"
)

type P []int

func (p P) Len() int           { return len(p) }
func (p P) Less(i, j int) bool { return p[i] < p[j] }
func (p P) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p *P) Push(x any)        { *p = append(*p, x.(int)) }
func (p *P) Pop() any {
	a := *p
	n := len(a) - 1
	x := a[n]
	*p = a[:n]
	return x
}

func main() {
	var (
		n, x, t int
		p       P
	)

	Scan(&n)
	for n > 0 {
		Scan(&x)
		Push(&p, x)
		n--
	}

	for p.Len() > 1 {
		n = Pop(&p).(int) + Pop(&p).(int)
		t += n
		Push(&p, n)
	}

	Printf("%.2f", float64(t)*0.05)
}