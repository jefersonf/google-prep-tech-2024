package module4

const (
	maxN = 10000
	maxK = 25
)

func RangeMinimumQuery(a []int, queries [][]int) []int {

	sparseTable := [maxN][maxK]int{}
	n := len(a)

	log := make([]int, maxN+1)
	log[1] = 0
	for i := 2; i <= maxN; i++ {
		log[i] = log[i/2] + 1
	}

	for i := range n {
		sparseTable[i][0] = a[i]
	}

	for j := 1; j < maxK; j++ {
		for i := 0; i+(1<<j) <= maxN; i++ {
			sparseTable[i][j] = min(sparseTable[i][j-1], sparseTable[i+(1<<(j-1))][j-1])
		}
	}

	answers := make([]int, 0)
	for _, q := range queries {
		l, r := q[0], q[1]
		j := log[r-l+1]
		answers = append(answers, min(sparseTable[l][j], sparseTable[r-(1<<j)+1][j]))

	}

	return answers
}
