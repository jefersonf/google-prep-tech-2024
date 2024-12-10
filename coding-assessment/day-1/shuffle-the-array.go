package day1

func shuffle(nums []int, n int) []int {
	shuffled := make([]int, 0, 2*n)
	for i := 0; i < n; i++ {
		shuffled = append(shuffled, nums[i], nums[n+i])
	}
	return shuffled
}
