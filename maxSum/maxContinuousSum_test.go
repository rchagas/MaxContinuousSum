package maxSum

import (
	"testing"
)

func TestGenerateSubList(t *testing.T) {
	t.Run("Teste Generate List start < end", func(t *testing.T) {
		start := 1
		end := 5
		list, error := GenerateSubList(start, end)
		wantedList := []int{1, 2, 3, 4, 5}

		if (list[0] != wantedList[0] ||
			len(list) != len(wantedList)) || error != nil {
			t.Errorf("result %v %v, want %v, dado %d %d", list, error, wantedList, start, end)
		}
	})

	t.Run("Teste Generate List start = end", func(t *testing.T) {
		start := 2
		end := 2
		list, error := GenerateSubList(start, end)
		wantedList := []int{2}

		if (list[0] != wantedList[0] ||
			len(list) != len(wantedList)) || error != nil {
			t.Errorf("result %v %v, want %v, dado %d %d", list, error, wantedList, start, end)
		}
	})

	t.Run("Teste Generate List start > end", func(t *testing.T) {
		start := 5
		end := 1
		list, error := GenerateSubList(start, end)
		wantedList := []int{}

		if error == nil {
			t.Errorf("result %v %v, want %v, dado %d %d", list, error, wantedList, start, end)
		}
	})
}

func TestMaxContinuousSum(t *testing.T) {

	t.Run("Teste soma", func(t *testing.T) {
		list := []int{1, 2, 3, 4, 5}

		sum, indexList := MaxContinuousSum(list)
		wantedSum := 15
		wantedList := []int{0, 1, 2, 3, 4}

		if sum != wantedSum ||
			indexList[0] != wantedList[0] ||
			len(indexList) != len(wantedList) {
			t.Errorf("result %d %v, want %d %v, dado %v", sum, indexList, wantedSum, wantedList, list)
		}
	})

	t.Run("Entrada: -2, 3, 5, -1, 4, -5", func(t *testing.T) {
		list := []int{-2, 3, 5, -1, 4, -5}

		sum, indexList := MaxContinuousSum(list)
		wantedSum := 11
		wantedList := []int{1, 2, 3, 4}

		if sum != wantedSum ||
			indexList[0] != wantedList[0] ||
			len(indexList) != len(wantedList) {
			t.Errorf("result %d %v, want %d %v, dado %v", sum, indexList, wantedSum, wantedList, list)
		}
	})

	t.Run("Entrada: -1000, -1, -1, -1, -1000", func(t *testing.T) {
		list := []int{-1000, -1, -1, -1, -1000}

		sum, indexList := MaxContinuousSum(list)
		wantedSum := -1
		wantedList := []int{1}

		if sum != wantedSum ||
			indexList[0] != wantedList[0] ||
			len(indexList)-1 != len(wantedList)-1 {
			t.Errorf("result %d %v, want %d %v, dado %v", sum, indexList, wantedSum, wantedList, list)
		}
	})

	t.Run("Entrada: 1, 2, 3, 4, -1000", func(t *testing.T) {
		list := []int{1, 2, 3, 4, -1000}

		sum, indexList := MaxContinuousSum(list)
		wantedSum := 10
		wantedList := []int{0, 1, 2, 3}

		if sum != wantedSum ||
			indexList[0] != wantedList[0] ||
			len(indexList)-1 != len(wantedList)-1 {
			t.Errorf("result %d %v, want %d %v, dado %v", sum, indexList, wantedSum, wantedList, list)
		}
	})

	t.Run("Entrada: -1000, 1, 2, 3, 4", func(t *testing.T) {
		list := []int{-1000, 1, 2, 3, 4}

		sum, indexList := MaxContinuousSum(list)
		wantedSum := 10
		wantedList := []int{1, 2, 3, 4}

		if sum != wantedSum ||
			indexList[0] != wantedList[0] ||
			len(indexList)-1 != len(wantedList)-1 {
			t.Errorf("result %d %v, want %d %v, dado %v", sum, indexList, wantedSum, wantedList, list)
		}
	})

	t.Run("Entrada: -1000, 2, -1, 2, -1000", func(t *testing.T) {
		list := []int{-1000, 2, -1, 2, -1000}

		sum, indexList := MaxContinuousSum(list)
		wantedSum := 3
		wantedList := []int{1, 2, 3}

		if sum != wantedSum ||
			indexList[0] != wantedList[0] ||
			len(indexList)-1 != len(wantedList)-1 {
			t.Errorf("result %d %v, want %d %v, dado %v", sum, indexList, wantedSum, wantedList, list)
		}
	})

}
