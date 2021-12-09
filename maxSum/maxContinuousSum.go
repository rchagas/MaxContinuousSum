package maxSum

import (
	"errors"
)

func MaxContinuousSum(listNumbers []int) (int, []int) {
	maxContinuousSum := listNumbers[0]
	startIndex := 0
	endIndex := 0
	nextSum := 0
	nextSumStartIndex := 0

	for i := 0; i < len(listNumbers); i++ {
		nextSum += listNumbers[i]
		if maxContinuousSum < nextSum {
			maxContinuousSum = nextSum
			startIndex = nextSumStartIndex
			endIndex = i
		}
		if nextSum < 0 {
			nextSum = 0
			nextSumStartIndex = i + 1
		}
	}
	listNumbers, _ = GenerateSubList(startIndex, endIndex)
	return maxContinuousSum, listNumbers
}

func GenerateSubList(startIndex int, endIndex int) ([]int, error) {
	if endIndex >= startIndex {
		list := make([]int, (endIndex - startIndex + 1))
		j := startIndex
		for i := 0; i < (endIndex - startIndex + 1); i++ {
			list[i] = j
			j++
		}
		return list, nil
	} else {
		return []int{}, errors.New("Start index bigger than end Index")
	}
}
