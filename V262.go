package main
import (
	. "container/heap"
	. "fmt"
)
func main() {
	var n, x, t int
	Scan(&n)
	p := make(PQ, 0, n)
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

type PQ []int

func (p PQ) Len() int           { return len(p) }
func (p PQ) Less(i, j int) bool { return p[i] < p[j] }
func (p PQ) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p *PQ) Push(x any)        { *p = append(*p, x.(int)) }
func (p *PQ) Pop() any {
	a := *p
	n := len(a) - 1
	x := a[n]
	*p = a[:n]
	return x
}