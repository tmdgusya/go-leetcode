package golangleetcode_test

import (
	"fmt"
	"sort"
	"testing"
)

// Space Complexity : O(N)
// Time Complexity: O(NlogN)
func TestCase1(t *testing.T) {
	arr1 := [11]int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}
	arr2 := [6]int{2, 1, 4, 3, 9, 6}

	expected := [11]int{2, 2, 2, 1, 4, 3, 3, 9, 6, 7, 19}

	result := relativeSortArray(arr1[:], arr2[:])

	fmt.Println(result)

	if !Equal(expected[:], result[:]) {
		t.Errorf("Error!")
	}
}

func relativeSortArray(arr1 []int, arr2 []int) []int {
	orderMap := saveOrderMap(arr1, arr2)

	fmt.Println(orderMap)

	sort.Slice(arr1, func(i, j int) bool {
		return orderMap[arr1[i]] < orderMap[arr1[j]]
	})

	fmt.Println(arr1)

	return arr1
}

func saveOrderMap(arr1 []int, arr2 []int) map[int]int {
	orderMap := make(map[int]int)

	for _, v := range arr1 {
		if orderMap[v] == 0 {
			orderMap[v] = indexOf(arr2, v)
		}
	}
	return orderMap
}

func indexOf(arr2 []int, target int) int {
	for i, v := range arr2 {
		if v == target {
			return i
		}
	}
	// maximum num
	return 1000 + target
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}
