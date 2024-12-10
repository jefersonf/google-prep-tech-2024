package sorting

func CountingSort(arr []int) []int {
	var n int
	for _, v := range arr { // O(n)
		n = max(n, v)
	}

	counter := make([]int, n+1)

	for _, v := range arr { // O(n)
		counter[v]++
	}

	for i := 1; i < len(counter); i++ { // O(k), sendo k = max(array)
		counter[i] += counter[i-1]
	}

	result := make([]int, len(arr))

	for j := len(arr) - 1; j >= 0; j-- { // O(n)
		amount := counter[arr[j]]
		result[amount-1] = arr[j]
		counter[arr[j]]--
	}

	return result
}
