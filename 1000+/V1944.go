package main
import . "fmt"
func main() {
	var (
		q                         [2e3][2e3]bool
		M, N, V, H, A, i, j, l, z int
		u                         = []int{0, 0, -1, 1, 0, 0}
		S                         = Scan
	)

	S(&M, &N, &V, &H)
	for j < V {
		S(&z)
		k := 0
		for k < N {
			k++
			q[k][z] = 1 > 0
		}
		j++
	}

	for l < H {
		S(&z)
		j = 0
		for j < M {
			j++
			q[z][j] = 1 > 0
		}
		l++
	}

	H = V*N + H*M - V*H

	for i < N {
		i++
		j = 0
		for j < M {
			j++
			if !q[i][j] {
				l = 0
				for l < 4 {
					z = i + u[l]
					V = j + u[l+2]
					if z > 0 && z <= N && V > 0 && V <= M && q[z][V] {
						A++
						break
					}
					l++
				}
			}
		}
	}

	Print(A, M*N-H-A, H)
}