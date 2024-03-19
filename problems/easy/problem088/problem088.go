package problem088

import (
	"sort"
)

const MaxLen = 200

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	if !isValidInput(nums1, nums2, m, n) {
		return nums1
	}

	return mergeArrays(nums1, nums2, m, n)
}

func mergeUsingSort(nums1 []int, m int, nums2 []int, n int) []int {
	if !isValidInput(nums1, nums2, m, n) {
		return nums1
	}

	for i := 0; i < len(nums2); i++ {
		nums1[m+i] = nums2[i]
	}

	sort.Ints(nums1)

	return nums1
}

func mergeArrays(nums1, nums2 []int, m, n int) []int {
	index1, index2 := m-1, n-1

	for i := m + n - 1; i >= 0; i-- {
		if index2 < 0 {
			break
		}

		if index1 >= 0 && nums1[index1] > nums2[index2] {
			nums1[i] = nums1[index1]
			index1--
		} else {
			nums1[i] = nums2[index2]
			index2--
		}
	}

	return nums1
}

func isValidInput(nums1, nums2 []int, m, n int) bool {
	if !isValidLength(m, n) || !isValidArrayLength(nums1, nums2, m, n) {
		return false
	}

	return true
}

func isValidLength(m, n int) bool {
	if m < 0 || m > MaxLen || n < 0 || n > MaxLen || m+n < 1 || m+n > MaxLen {
		return false
	}

	return true
}

func isValidArrayLength(nums1, nums2 []int, m, n int) bool {
	if len(nums1) != m+n || len(nums2) != n {
		return false
	}

	return true
}
