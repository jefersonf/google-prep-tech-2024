package module4

func LongestPalindromicSubsequence(s string) int {

	var lps func(l, r int) int

	lps = func(l, r int) int {
		if l > r {
			return 0
		}
		if l == r {
			return 1
		}
		if s[l] == s[r] {
			return 2 + lps(l+1, r-1)
		}

		return max(lps(l+1, r), lps(l, r-1))
	}

	return lps(0, len(s)-1)
}
