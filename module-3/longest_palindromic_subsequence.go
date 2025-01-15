package module3

func LongestPalindromicSubsequence(s string) int {

	N := len(s)
	memo := make([]int, N*N+1)

	var lps func(l, r int) int
	lps = func(l, r int) int {
		if l > r {
			return 0
		}

		q := l*N + r
		if memo[q] != 0 {
			return memo[q]
		}

		if l == r {
			memo[q] = 1
			return memo[q]
		}

		if s[l] == s[r] {
			memo[q] = 2 + lps(l+1, r-1)
		} else {
			memo[q] = max(lps(l+1, r), lps(l, r-1))
		}

		return memo[q]
	}

	return lps(0, len(s)-1)
}
