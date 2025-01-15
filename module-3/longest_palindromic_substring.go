package module3

func LongestPalindromicSubstring(s string) string {

	n := len(s)
	memo := make([]bool, n*n)

	for i := range n {
		memo[i*n+1] = true
	}

	l, r := 0, 0
	for i := range n - 1 {
		if s[i] == s[i+1] {
			memo[i*n+i+1] = true
			l, r = i, i+1
		}
	}

	for k := 2; k < n; k++ {
		for i := range n - k {
			j := i + k
			if s[i] == s[j] && memo[(i+1)*n+j-1] {
				memo[i*n+j] = true
				l, r = i, j
			}
		}
	}

	return s[l : r+1]
}
