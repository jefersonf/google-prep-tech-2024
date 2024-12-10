package sorting

import (
	"fmt"
	"preptech24/module-1/queue"
)

type radixLevel [10]queue.Queue

type radix struct {
	array     []int
	nextLevel *radixLevel
}

func RadixSort(arr []int) []int {
	if len(arr) == 0 {
		return make([]int, 0)
	}
	minValue := getMinValue(arr)

	nextLevelValues := make([]int, len(arr))
	var nextLevel radixLevel
	for i, v := range arr {
		nextLevel[v%10].Enqueue(v)
		nextLevelValues[i] = v / 10
	}

	//r := &radix{array: nextLevelValues, nextLevel: &nextLevel}

	fmt.Println("minVal = ", minValue)

	return []int{}
}

func getMinValue(a []int) int {
	minVal := 1 << 20
	for _, v := range a {
		minVal = min(minVal, v)
	}
	return minVal
}
