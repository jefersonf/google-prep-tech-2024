package day2

func linearSearch(nums []int) int {
	var positives, negatives int
	for _, value := range nums {
		if value < 0 {
			positives += 1
		} else if value > 0 {
			negatives += 1
		}
	}
	return max(positives, negatives)
}

func linearSearch2(nums []int) int {
	var positives, negatives = 0, 0
	for _, v := range nums {
		if v >= 0 {
			break
		}
		negatives += 1
	}

	for i := len(nums) - 1; i >= 0; i-- {
		v := nums[i]
		if v <= 0 {
			break
		}
		positives += 1
	}

	return max(negatives, positives)

}

func binarySearch(nums []int) int {
	countNegatives := func(a []int) int {
		if a[len(a)-1] < 0 {
			return len(a)
		}
		if a[0] >= 0 {
			return 0
		}
		var left, mid, right = 0, -1, len(a) - 1
		for left <= right {
			mid = (left + right) / 2
			if a[mid] < 0 && a[mid+1] >= 0 {
				break
			}
			if a[mid] < 0 {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		return mid + 1
	}

	countPositives := func(a []int) int {
		if a[0] > 0 {
			return len(a)
		}
		if a[len(a)-1] <= 0 {
			return 0
		}
		var left, mid, right = 0, -1, len(a) - 1
		for left <= right {
			mid = (left + right) / 2
			if a[mid] > 0 && a[mid-1] <= 0 {
				break
			}
			if a[mid] > 0 {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		return len(a) - mid
	}

	negatives := countNegatives(nums)
	positives := countPositives(nums)

	return max(positives, negatives)
}

func maximumCount(nums []int) int {
	return binarySearch(nums)
}
