package module4

const (
	maxN = 10000
	maxK = 25
)

func RMQ(a []int, queries [][]int) []int {

	st := [maxN][maxK]int{}
	n := len(a)

	log := make([]int, maxN+1)
	log[1] = 0
	for i := 2; i <= maxN; i++ {
		log[i] = log[i/2] + 1
	}

	for i := range n {
		st[i][0] = a[i]
	}

	for j := 1; j < maxK; j++ {
		for i := 0; i+(1<<j) <= maxN; i++ {
			st[i][j] = min(st[i][j-1], st[i+(1<<(j-1))][j-1])
		}
	}

	answers := make([]int, 0)
	for _, q := range queries {
		l, r := q[0], q[1]
		j := log[r-l+1]
		answers = append(answers, min(st[l][j], st[r-(1<<j)+1][j]))

	}

	return answers
}
