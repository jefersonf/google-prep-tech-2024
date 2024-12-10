package day2

func search(nums []int, target int) int {
	var (
		left  = 0
		right = len(nums) - 1
	)
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}